package ticket

type TicketForInsert struct {
	title       string
	description string
	creator_id  int
}

type TicketRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}