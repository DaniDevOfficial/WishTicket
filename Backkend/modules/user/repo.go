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

func CreateUserInDB(userData DBNewUser, db *sql.DB) error {
	sql := "INSERT INTO user (username, email, passwordHash) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(sql)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(userData.username, userData.email, userData.passwordHash)
	return err
}
