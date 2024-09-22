package ticket

import (
	"database/sql"
	"net/http"
)

func RegisterTicketRoute(router *http.ServeMux, db *sql.DB) {
	baseRoute := "/ticket"
	router.HandleFunc(baseRoute, func(w http.ResponseWriter, r *http.Request) {
		handleTicket(w, r, db)
	})
	router.HandleFunc(baseRoute+"/assignee", func(w http.ResponseWriter, r *http.Request) {
		handleAssignee(w, r, db)
	})
	router.HandleFunc(baseRoute+"/status", func(w http.ResponseWriter, r *http.Request) {
		handleStatus(w, r, db)
	})

	router.HandleFunc(baseRoute+"/assignee/tickets", func(w http.ResponseWriter, r *http.Request) {
		handleTicketAssignee(w, r, db)
	})

}

func handleTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		CreateNewTicket(w, r, db)
	}
	if r.Method == http.MethodGet {
		GetAllOwnedTickets(w, r, db)
	}
}

func handleAssignee(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		AddAssigneeToTicket(w, r, db)
	}
}

func handleStatus(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPut {
		ChangeTicketStatus(w, r, db)
	}
}

func handleTicketAssignee(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	if r.Method == http.MethodGet {
		GetAssignedTickets(w, r, db)
	}
}
