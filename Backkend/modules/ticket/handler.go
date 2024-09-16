package ticket

import (
	"database/sql"
	"fmt"
	"net/http"
)

func RegisterTicketRoute(router *http.ServeMux, db *sql.DB) {

	router.HandleFunc("/ticket", func(w http.ResponseWriter, r *http.Request) {
		handleTicket(w, r, db)
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "wasd")
	})
}

func handleTicket(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		CreateNewTicket(w, r, db)
	} else {
		fmt.Fprintf(w, "Ticket route")
	}
}
 