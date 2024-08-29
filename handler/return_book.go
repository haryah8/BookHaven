package handler

import (
	"BookHaven/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	LateFeePerDay     = 1000
	MaxDaysWithoutFee = 5
)

func ReturnBook(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request payload (ISBN only)
		var request models.ReturnBookDto // Assuming it has an `ISBN` field
		if err := c.Bind(&request); err != nil {
			logger.Error("Error Binding Request")
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
		}

		// Get user claims from context
		user := c.Get("user")
		claims, ok := user.(*models.Claims) // Type assertion
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user claims"})
		}

		// Fetch borrowing details using ISBN and user_id
		var borrowingID, bookId int
		var borrowedAtStr string // Store as string

		err := db.QueryRow(`
			SELECT b.id, br.id, br.borrowed_at 
			FROM borrowings br 
			JOIN books b ON br.book_id = b.id 
			WHERE b.isbn = ? AND br.user_id = ? AND br.status = 'borrowed'`,
			request.ISBN, claims.UserId).Scan(&bookId, &borrowingID, &borrowedAtStr)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Borrowing record not found for this ISBN"})
			}
			logger.Error("Error querying borrowing record: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking borrowing record"})
		}

		// Parse borrowedAtStr to time.Time using the appropriate format
		const layout = "2006-01-02 15:04:05"
		borrowDate, err := time.Parse(layout, borrowedAtStr)
		if err != nil {
			logger.Error("Error converting borrow date: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error converting borrow date"})
		}

		// Calculate late fee if the book is returned after 5 days
		daysBorrowed := int(time.Since(borrowDate).Hours() / 24)
		var lateFee int
		if daysBorrowed > MaxDaysWithoutFee {
			lateFee = (daysBorrowed - MaxDaysWithoutFee) * LateFeePerDay
		}

		// Check user's balance if there's a late fee
		if lateFee > 0 {
			var balance int
			err := db.QueryRow(`SELECT balance FROM users WHERE id = ?`, claims.UserId).Scan(&balance)
			if err != nil {
				logger.Error("Error querying user balance: ", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking user balance"})
			}

			// Ensure the user has enough balance to cover the late fee
			if balance < lateFee {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Insufficient balance. Please top up your account to cover the late fee of " + strconv.Itoa(lateFee) + " before returning the book.",
				})
			}

			// Deduct late fee from user's balance
			_, err = db.Exec(`UPDATE users SET balance = balance - ? WHERE id = ?`, lateFee, claims.UserId)
			if err != nil {
				logger.Error("Error deducting late fee from balance: ", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deducting late fee"})
			}
		}

		// Update borrowing status to 'returned'
		_, err = db.Exec(`UPDATE borrowings SET status = 'returned', returned_at = ? WHERE id = ?`, time.Now(), borrowingID)
		if err != nil {
			logger.Error("Error updating borrowing status: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error updating borrowing status"})
		}

		// Increase the available copies of the book
		_, err = db.Exec(`UPDATE books SET available_copies = available_copies + 1 WHERE id = ?`, bookId)
		if err != nil {
			logger.Error("Error updating book availability: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error updating book availability"})
		}
		if lateFee > 0 {
			return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Book returned successfully with additional late fee %d", lateFee)})

		} else {
			return c.JSON(http.StatusOK, map[string]string{"message": "Book returned successfully"})

		}
	}
}
