package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wishticket/util/hashing"
	"wishticket/util/jwt"
)

func CreateNewUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("Users called")

	var newUser RequestNewUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	log.Println("Users called")

	if err != nil {
		log.Println(err)
		return
	}

	userId, err := GetUserIdByName(newUser.Username, db)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return
	}
	log.Println("Users called")

	if userId != -1 {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User already exists")
		return
	}

	hashedPassword, err := hashing.HashPassword(newUser.Password)
	if err != nil {
		log.Println(err)
		return
	}

	userInDB := DBNewUser{
		username:      newUser.Username,
		email:         newUser.Email,
		password_hash: hashedPassword,
	}

	fmt.Println(userInDB)
	id, err := CreateUserInDB(userInDB, db)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	jwtUserData := jwt.JWTUser{
		Username: userInDB.username,
		UserId:   id,
	}
	jwtToken, err := jwt.CreateToken(jwtUserData)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(map[string]string{"token": jwtToken})
	if err != nil {
		return
	}
}

func SignIn(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var credentials SignInCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println(err)
		return
	}
	userData, err := GetUserByName(credentials.Username, db)
	//TODO: Implement Data validation, so all required data is present
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(userData.password_hash)
	if !hashing.CheckHashedString(userData.password_hash, credentials.Password) {
		fmt.Fprintf(w, "Wrong username or Password")
		return
	}
	jwtUserData := jwt.JWTUser{
		Username: userData.username,
		UserId:   userData.user_id,
	}
	token, err := jwt.CreateToken(jwtUserData)
	if err != nil {
		fmt.Fprintf(w, "Error happened")
		log.Println(err)
		return
	}

	fmt.Fprintf(w, token)
}
