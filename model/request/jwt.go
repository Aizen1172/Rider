package request

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	ID uint
	Name string `json:"name"`
	jwt.StandardClaims
}