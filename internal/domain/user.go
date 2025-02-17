package domain

type User struct {
	ID   int64
	Info UserInfo
}

type UserInfo struct {
	Email    string
	Username string
	Password string
}
