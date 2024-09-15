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

type DBNewUser struct {
	username     string
	email        string
	passwordHash string
}

type DBUserData struct {
	user_id      int
	username     string
	email        string
	passwordHash string
}

type SignInCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateNewUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newUser RequestNewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		log.Println(err)
		return
	}
	_, err = GetUserIdByName(newUser.Username, db)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		log.Println(err)
		return
	}

	userInDB := DBNewUser{
		username:     newUser.Username,
		email:        newUser.Email,
		passwordHash: hashedPassword,
	}
	fmt.Println(userInDB)
	CreateUserInDB(userInDB, db)
}

func SignIn(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var credentials SignInCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println(err)
		return
	}

	//TODO: Implement Data validation, so all required data is present
	passwordHash, err := GetUserPasswordHashByName(credentials.Username, db)
	if err != nil {
		log.Println(err)
		return
	}

	if !util.CheckHashedString(passwordHash, credentials.Password) {
		fmt.Fprintf(w, "Wrong username or Password")
		return
	}
	fmt.Fprintf(w, "yayyy your logged in 😊😊😊")

}
