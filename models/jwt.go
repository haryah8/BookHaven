package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserId int
	Email  string
	jwt.StandardClaims
}
