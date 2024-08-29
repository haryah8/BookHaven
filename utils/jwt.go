package utils

import (
	"BookHaven/config"
	"BookHaven/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// Secret key used for signing tokens
var jwtKey = []byte(config.SECRET_KEY) // Change this to a secure key

// Claims struct to define the payload of the JWT

// GenerateToken generates a JWT token for a given user ID
func GenerateToken(email string, id int) (string, error) {
	claims := models.Claims{
		UserId: id,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "harya",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken parses and validates the JWT token
func ValidateToken(tokenStr string) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		fmt.Println(claims.UserId)
		fmt.Println(claims.Email)
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
