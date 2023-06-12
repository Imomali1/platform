package auth

const (
	getUser = `SELECT * FROM users WHERE username = $1 AND password = $2`
)
