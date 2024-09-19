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

type StatusRequest struct {
	TicketId int    `json:"ticketId"`
	Status   string `json:"status"`
}

type AddAssigneeRequest struct {
	TicketId   int `json:"ticketId"`
	AssignedId int `json:"assignedId"`
}

type TicketFromDB struct {
	TicketId    int            `json:"ticketId"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	CreatorId   int            `json:"creatorId"`
	Status      string         `json:"status"`
}
