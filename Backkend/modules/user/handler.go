package user

import (
	"database/sql"
	"net/http"
)

func RegisterUserRoute(router http.ServerMux, db *sql.DB) {

	router.handleRequest("/users", func(w http.ResponseWriter, r *http.Request) {
		handleUsers(w, r, db)
	})
}

func handleUsers(w http.ResponseWriter, r *http.Request, db *sql.DB)
