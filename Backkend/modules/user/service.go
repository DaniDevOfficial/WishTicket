package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wishticket/util"
)

type RequestNewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateNewUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newUser RequestNewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		log.Println(err)
		return
	}

	hashedPassword, err := util.HashPassword("wasd")
	fmt.Fprintf(w, hashedPassword)
	fmt.Println(hashedPassword)
	fmt.Println(util.CheckHashedString(hashedPassword, "wasd"))

}
