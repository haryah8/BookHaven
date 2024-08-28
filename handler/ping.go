package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

// PingMe checks the database connection and responds with a success message
func PingMe(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ping the database
		if err := db.Ping(); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Database connection failed")
		}

		// Respond with success
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "Success",
			"message": "Ping!",
		})
	}
}
