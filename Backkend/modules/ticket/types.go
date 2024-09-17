package ticket

import (
	"database/sql"
)

type TicketForInsert struct {
	title       string
	description string
	creator_id  int
}

type TicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AddAssigneeRequest struct {
	TicketId   int `json:"ticketId"`
	AssignedId int `json:"assignedId"`
}

type TicketFromDB struct {
	ticket_id   int
	title       string
	description sql.NullString
	creator_id  int
}
