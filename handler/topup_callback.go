package handler

import (
	"BookHaven/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func TopUpCallback(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request models.TopUpCallbackRequestDto
		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		}

		// Query to get existing transaction details
		var userId int
		var amount int
		var status string
		err := db.QueryRow(`SELECT user_id, amount, status FROM transactions WHERE transaction_id = ?`, request.ExternalId).Scan(&userId, &amount, &status)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "Error",
					"message": "Data not found",
				})
			}
			logger.Error("Error querying database: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":  "Error",
				"message": "Internal Server Error",
			})
		}

		var updateStatus = "pending"
		if request.Status == "PAID" || request.Status == "COMPLETED" {
			updateStatus = "completed"
		} else if request.Status == "FAILED" {
			updateStatus = "failed"
		}

		// Check if the callback status matches the transaction
		if amount == request.Amount && status == "pending" {
			// Update transaction status
			_, err = db.Exec(`UPDATE transactions SET status = ? WHERE transaction_id = ?`, updateStatus, request.ExternalId)
			if err != nil {
				logger.Error("Error updating transaction status: ", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"status": "Error", "message": "Failed to update transaction"})
			}

			// Add balance to user if the status is "completed"
			if updateStatus == "completed" {
				_, err = db.Exec(`UPDATE users SET balance = balance + ? WHERE id = ?`, request.Amount, userId)
				if err != nil {
					logger.Error("Error updating user balance: ", err)
					return c.JSON(http.StatusInternalServerError, map[string]string{"status": "Error", "message": "Failed to update user balance"})
				}
			}

			return c.JSON(http.StatusOK, map[string]string{"status": "Success", "message": "Top-up processed successfully"})
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"status": "Error", "message": "Invalid amount or status"})
		}
	}
}
