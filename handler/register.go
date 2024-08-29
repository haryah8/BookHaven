package handler

import (
	"BookHaven/models"
	"BookHaven/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and name
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.RegisterDto true "User Registration"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.RegisterDto // Assuming you have a User model
		if err := c.Bind(&req); err != nil {
			logger.Error("Error binding request body: ", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"status":  "Error",
				"message": "Invalid input",
			})
		}

		fmt.Println(req) // DEBUGGING --------------------------------------------

		// Hash the password
		hashedPassword, err := utils.HashPassword(req.Password)
		fmt.Printf("%s, %s", req.Password, hashedPassword) // DEBUGGING --------------------------------------------
		if err != nil {
			logger.Error("Error hashing password: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "Error",
				"message": "Failed to hash password",
			})
		}

		// Insert user into database
		_, err = db.Exec(`INSERT INTO users (email, password_hash, name) VALUES (?, ?, ?)`,
			req.Email, hashedPassword, req.Name)
		if err != nil {
			logger.Error("Error inserting user into database: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "Error",
				"message": "Failed to register user",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status":  "Success",
			"message": "User registered successfully",
		})
	}
}
