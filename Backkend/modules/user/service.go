package user

import (
	"database/sql"
	"fmt"
	"net/http"
)

type RequestNewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateNewUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Fprintf(w, "hello world")
}
