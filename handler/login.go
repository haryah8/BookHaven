package handler

import (
	"BookHaven/models"
	"BookHaven/utils" // Utility package for password hashing and JWT
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Login handles user login
func Login(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.LoginDto // Assuming you have a User model

		if err := c.Bind(&req); err != nil {
			logger.Error("Error binding request body: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"status":  "Error",
				"message": "Invalid input",
			})
		}

		var storedPasswordHash string
		var userId int
		err := db.QueryRow(`SELECT password_hash, id FROM users WHERE email = ?`, req.Email).Scan(&storedPasswordHash, &userId)
		if err != nil {
			if err == sql.ErrNoRows {

				return c.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "Error",
					"message": "Invalid credentials",
				})
			}
			logger.Error("Error querying database: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "Error",
				"message": "Failed to authenticate user",
			})
		}

		fmt.Println(storedPasswordHash) // DEBUGGING --------------------------------------------
		fmt.Println(req.Password)       // DEBUGGING --------------------------------------------
		// Check password
		if utils.CheckPasswordHash(storedPasswordHash, req.Password) {
			fmt.Printf("%s, %s, %s", storedPasswordHash, req.Password, err) // DEBUGGING --------------------------------------------
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"status":  "Error",
				"message": "Invalid credentials",
			})
		}

		fmt.Println(userId) // DEBUGGING --------------------------------------------
		// Generate JWT token
		token, err := utils.GenerateToken(req.Email, userId)
		if err != nil {
			logger.Error("Error generating token: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "Error",
				"message": "Failed to generate token",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status":  "Success",
			"message": "Login successful",
			"token":   token,
		})
	}
}
