package ticket

import (
	"database/sql"
)

// Ticket

func GetAllOwnedTicketsFromDB(userId int, db *sql.DB) ([]TicketFromDB, error) {
	sql := "SELECT ticket_id, title, description, creator_id FROM ticket WHERE creator_id = ?"
	var tickets []TicketFromDB
	rows, err := db.Query(sql, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ticket TicketFromDB
		err := rows.Scan(&ticket.TicketId, &ticket.Title, &ticket.Description, &ticket.CreatorId)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil

}

func CreateNewTicketInDB(ticketData TicketForInsert, db *sql.DB) error {

	tx, err := db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	lastId, err := insertNewTicket(ticketData, tx)

	if err != nil {
		return err
	}

	err = insertNewTicketStatus(int(lastId), tx)

	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func insertNewTicket(ticketData TicketForInsert, tx *sql.Tx) (int, error) {
	sql := "INSERT INTO ticket (title, description, creator_id) VALUES (?, ?, ?)"
	stmt, err := tx.Prepare(sql)

	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ticketData.title, ticketData.description, ticketData.creator_id)

	if err != nil {
		return -1, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return int(lastId), nil
}

func GetTicketById(ticketId int, db *sql.DB) (*TicketFromDB, error) {

	sql := "SELECT * FROM ticket WHERE ticket_id = ?"

	row := db.QueryRow(sql, ticketId)
	var ticket TicketFromDB
	err := row.Scan(&ticket.TicketId, &ticket.Title, &ticket.Description, &ticket.CreatorId)
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

	_, err = stmt.Exec(ticket_id, "Open")
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
