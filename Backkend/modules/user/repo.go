package user

import (
	"database/sql"
)

func GetUserIdByName(username string, db *sql.DB) (int, error) {
	sql := "SELECT user_id FROM user WHERE username = ?"
	row := db.QueryRow(sql, username)
	var userId int
	err := row.Scan(&userId)
	return userId, err
}

func GetUserPasswordHashByName(username string, db *sql.DB) (string, error) {
	sql := "SELECT passwordHash FROM user WHERE username = ?"
	row := db.QueryRow(sql, username)
	var passwordHash string
	err := row.Scan(&passwordHash)
	return passwordHash, err
}

func CreateUserInDB(userData DBNewUser, db *sql.DB) error {
	sql := "INSERT INTO user (username, email, passwordHash) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(userData.username, userData.email, userData.passwordHash)
	return err
}

type UserFromDB struct {
	user_id      int
	username     string
	email        string
	passwordHash string
}

func GetUserById(id int, db *sql.DB) (UserFromDB, error) {

	sql := "SELECT * FROM user WHERE user_id = ?"
	row := db.QueryRow(sql, id)

	var userData UserFromDB
	err := row.Scan(&userData)
	return userData, err

}
