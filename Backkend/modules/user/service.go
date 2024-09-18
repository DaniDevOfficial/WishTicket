package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wishticket/util/hashing"
)

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

	hashedPassword, err := hashing.HashPassword(newUser.Password)
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

	if !hashing.CheckHashedString(passwordHash, credentials.Password) {
		fmt.Fprintf(w, "Wrong username or Password")
		return
	}
	fmt.Fprintf(w, "yayyy your logged in 😊😊😊")
	// TODO: Should return a JWT
}
