package handler

import (
	"BookHaven/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetBorrowedBooks(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get user claims from context
		user := c.Get("user")
		claims, ok := user.(*models.Claims) // Type assertion
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user claims"})
		}

		// Query to get the books that the user is currently borrowing
		rows, err := db.Query(`
			SELECT b.title, b.author, b.isbn, br.borrowed_at 
			FROM borrowings br
			JOIN books b ON br.book_id = b.id
			WHERE br.user_id = ? AND br.status = 'borrowed'`,
			claims.UserId)
		if err != nil {
			logger.Error("Error querying borrowed books: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving borrowed books"})
		}
		defer rows.Close()

		// Prepare the response with the list of borrowed books
		var borrowedBooks []models.BorrowedBookResponse
		for rows.Next() {
			var book models.BorrowedBookResponse
			var borrowedAt string
			if err := rows.Scan(&book.Title, &book.Author, &book.ISBN, &borrowedAt); err != nil {
				logger.Error("Error scanning borrowed books: ", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error processing borrowed books"})
			}

			// Parse borrowedAt string to time.Time
			const layout = "2006-01-02 15:04:05"
			borrowDate, err := time.Parse(layout, borrowedAt)
			if err != nil {
				logger.Error("Error converting borrow date: ", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error converting borrow date"})
			}
			book.BorrowedAt = borrowDate

			borrowedBooks = append(borrowedBooks, book)
		}

		// Check for errors from iterating over rows
		if err := rows.Err(); err != nil {
			logger.Error("Error iterating over borrowed books rows: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving borrowed books"})
		}

		// Return the list of borrowed books
		return c.JSON(http.StatusOK, borrowedBooks)
	}
}
