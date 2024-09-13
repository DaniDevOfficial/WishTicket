package user

import (
	"database/sql"
	"fmt"
	"net/http"
)

func RegisterUserRoute(router *http.ServeMux, db *sql.DB) {

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handleUsers(w, r, db)
	})
}

func handleUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		CreateNewUser(w, r, db)
	}
	fmt.Fprintf(w, "wasd")
}
