package ticket

import (
	"database/sql"
	"log"
)

// Ticket

func GetAllOwnedTicketsFromDB(userId int, requesterId int, db *sql.DB) ([]TicketFromDB, error) {
	sql := `
	SELECT DISTINCT 
		t.ticket_id, 
		t.title, 
		t.description, 
		t.visibility,
		t.creator_id,
		ts.status
	FROM 
		ticket t
	JOIN
		ticket_status ts ON t.ticket_id = ts.ticket_id AND t.creator_id = ?
	LEFT JOIN
		ticket_assigned ta ON t.ticket_id = ta.ticket_id 
	WHERE
		t.visibility = 'PUBLIC'
		OR ta.assigned_id = ?
		OR t.creator_id = ?
`

	var tickets []TicketFromDB
	rows, err := db.Query(sql, userId, requesterId, requesterId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ticket TicketFromDB
		err := rows.Scan(&ticket.TicketId, &ticket.Title, &ticket.Description, &ticket.Visibility, &ticket.CreatorId, &ticket.Status)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil

}

func GetAssignedTicketsFromDB(userId int, requesterId int, db *sql.DB) ([]TicketFromDB, error) {

	sql := `
			SELECT DISTINCT
				t.ticket_id,
				t.title,
				t.description,
				t.visibility,
				t.creator_id,
				ts.status
			FROM
				ticket t
			JOIN
				ticket_assigned ta ON ta.ticket_id = t.ticket_id
			JOIN
				ticket_status ts ON t.ticket_id = ts.ticket_id
			WHERE
				t.visibility = 'PUBLIC' 
				OR ta.assigned_id = ? 
				OR t.creator_id = ?
`

	var tickets []TicketFromDB
	rows, err := db.Query(sql, userId, requesterId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ticket TicketFromDB
		err := rows.Scan(&ticket.TicketId, &ticket.Title, &ticket.Description, &ticket.Visibility, &ticket.CreatorId, &ticket.Status)
		if err != nil {
			return nil, err
		}
		if ticket.Description.Valid {
			ticket.Description = ticket.Description
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func CreateNewTicketInDB(ticketData TicketForInsert, db *sql.DB) (int, error) {

	tx, err := db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	lastId, err := insertNewTicket(ticketData, tx)
	if err != nil {
		return -1, err
	}

	err = insertNewTicketStatus(lastId, tx)

	if err != nil {
		return -1, err
	}
	if err := tx.Commit(); err != nil {
		return -1, err
	}

	return lastId, nil
}

func insertNewTicket(ticketData TicketForInsert, tx *sql.Tx) (int, error) {
	log.Println(ticketData)
	sqlString := `	INSERT INTO 
					    ticket 
					    (title, description, visibility, dueDate, creator_id) 
					VALUES 
					    (?, ?, ?, ?, ?)`
	stmt, err := tx.Prepare(sqlString)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ticketData.title, ticketData.description, ticketData.visibility, ticketData.dueDate, ticketData.creator_id)

	if err != nil {
		return -1, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return int(lastId), nil
}

func GetTicketFromDB(ticketId int, requesterId int, db *sql.DB) (*TicketFromDB, error) {

	sql := `
		SELECT 
			t.ticket_id, 
			t.title, 
			t.description,
			t.dueDate,
			t.visibility,
			t.creator_id,
			ts.status
		FROM
			ticket t
		JOIN
			ticket_status ts ON t.ticket_id = ts.ticket_id
		LEFT JOIN 
			ticket_assigned ta ON t.ticket_id = ta.ticket_id
		WHERE
			t.ticket_id = ? 
        AND (
            t.visibility = 'PUBLIC' 
            OR ta.assigned_id = ?
            OR t.creator_id = ?
            )
	`
	row := db.QueryRow(sql, ticketId, requesterId, requesterId)
	var ticket TicketFromDB
	err := row.Scan(&ticket.TicketId, &ticket.Title, &ticket.Description, &ticket.DueDate, &ticket.Visibility, &ticket.CreatorId, &ticket.Status)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

// Status

func insertNewTicketStatus(ticket_id int, tx *sql.Tx) error {
	sql := "INSERT INTO ticket_status (ticket_id, status) VALUES (?, ?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ticket_id, "asdf")
	if err != nil {
		return err
	}
	return nil
}

func UpdateTicketStatus(newStatus StatusRequest, db *sql.DB) (sql.Result, error) {

	sql := "UPDATE ticket_status SET status = ? WHERE ticket_id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newStatus.Status, newStatus.TicketId)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Comments

// assignee

func CreateNewAssignment(newAssignment AddAssigneeRequest, db *sql.DB) error {
	sql := "INSERT INTO ticket_assigned (ticket_id, assigned_id) VALUES (?, ?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(newAssignment.TicketId, newAssignment.AssignedId)
	if err != nil {
		return err
	}
	return nil
}

// Other
