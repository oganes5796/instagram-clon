package model

type User struct {
	ID   int64    `db:"id"`
	Info UserInfo `db:""`
}

type UserInfo struct {
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string `db:"password"`
}
