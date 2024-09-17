package ticket

import (
	"database/sql"
	"net/http"
)

func RegisterTicketRoute(router *http.ServeMux, db *sql.DB) {

	router.HandleFunc("/ticket", func(w http.ResponseWriter, r *http.Request) {
		handleTicket(w, r, db)
	})
	router.HandleFunc("/ticket/assignee", func(w http.ResponseWriter, r *http.Request) {
		handleAssignee(w, r, db)
	})

}

func handleTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		CreateNewTicket(w, r, db)
	}
}

func handleAssignee(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		AddAssigneeToTicket(w, r, db)
	}
}
