package user

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

type SignInCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserFromDB struct {
	user_id      int
	username     string
	email        string
	passwordHash string
}
