package ticket

import (
	"database/sql"
)

const (
	VisibilityPublic  = "PUBLIC"
	VisibilityPrivate = "PRIVATE"
)

type TicketForInsert struct {
	title       string
	description string
	visibility  string
	creator_id  int
}

type TicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
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
	Visibility  string         `json:"visibility"`
	CreatorId   int            `json:"creatorId"`
	Status      string         `json:"status"`
}
