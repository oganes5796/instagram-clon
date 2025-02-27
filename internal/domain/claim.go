package domain

import "github.com/dgrijalva/jwt-go"

const (
	ExamplePath = "/user_v1.UserV1/Get"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}
