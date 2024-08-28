package handler

import (
	"BookHaven/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// PingMe checks the database connection and responds with a success message
func PingMe(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ping the database
		if err := db.Ping(); err != nil {
			response := c.JSON(http.StatusOK, map[string]string{
				"status":  "Error",
				"message": err.Error(),
			})
			return response
		}

		response := c.JSON(http.StatusOK, map[string]string{
			"status":  "Success",
			"message": "Ping!",
		})

		// Respond with success
		return response
	}
}

// ProtectedEndpoint is a protected API endpoint
func ProtectedEndpoint(c echo.Context) error {
	// Get the user claims from context
	claims := c.Get("user").(*models.Claims)
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "Success",
		"message": "Access granted",
		"user":    claims.Email,
	})
}
