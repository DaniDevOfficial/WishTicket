package ticket

import (
	"database/sql"
)

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
