package handler

import (
	"BookHaven/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetUserBalance(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get user claims from context
		user := c.Get("user")
		claims, ok := user.(*models.Claims) // Type assertion
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user claims"})
		}

		// Query to get the user's balance
		var balance int
		err := db.QueryRow(`SELECT balance FROM users WHERE id = ?`, claims.UserId).Scan(&balance)
		if err != nil {
			if err == sql.ErrNoRows {
				// User not found
				return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
			}
			// Handle other SQL errors
			logger.Error("Error querying user balance: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving balance"})
		}

		// Return the user's balance
		return c.JSON(http.StatusOK, map[string]int{"balance": balance})
	}
}
