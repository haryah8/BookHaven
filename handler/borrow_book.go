package handler

import (
	"BookHaven/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// BorrowBook handles the borrowing of a book by an authenticated user.
// @Summary Borrow a Book
// @Description Allows an authenticated user to borrow a book. The user must have a balance of at least 50,000 and cannot exceed the borrowing limit of 3 books. Handles errors such as insufficient balance, exceeding borrowing limit, or book availability issues.
// @Tags User
// @Accept json
// @Produce json
// @Param body body models.BorrowBookDto true "Borrow Book Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 403 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /user/book/borrow [post]
func BorrowBook(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request payload
		var request models.BorrowBookDto
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

		// Check user's balance
		var balance int
		err := db.QueryRow(`SELECT balance FROM users WHERE id = ?`, claims.UserId).Scan(&balance)
		if err != nil {
			logger.Error("Error querying user balance: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking user balance"})
		}

		// Ensure the user has a balance of at least 50,000
		if balance < 50000 {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "Insufficient balance. Please top up your account to at least 50,000 before borrowing a book."})
		}

		// Check if the user can borrow more books
		var borrowedCount int
		err = db.QueryRow(`SELECT COUNT(*) FROM borrowings WHERE user_id = ? AND status = 'borrowed'`, claims.UserId).Scan(&borrowedCount)
		if err != nil {
			logger.Error("Error querying borrowings: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking borrowings"})
		}

		// Borrowing limit: User cannot borrow more than 3 books
		if borrowedCount >= 3 {
			return c.JSON(http.StatusForbidden, map[string]string{"message": "You have reached the borrowing limit of 3 books."})
		}

		// Check book availability
		var availableCopies int
		var bookId int
		err = db.QueryRow(`SELECT id, available_copies FROM books WHERE isbn = ?`, request.ISBN).Scan(&bookId, &availableCopies)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ISBN"})
			}
			logger.Error("Error querying book data: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking book availability"})
		}

		// Ensure the book has available copies
		if availableCopies < 1 {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "No copies available for borrowing"})
		}

		// Insert new borrowing record
		_, err = db.Exec(`INSERT INTO borrowings (user_id, book_id, status) VALUES (?, ?, ?)`, claims.UserId, bookId, "borrowed")
		if err != nil {
			logger.Error("Error inserting borrowing data into database: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create borrowing record"})
		}

		// Deduct available copies of the book
		_, err = db.Exec(`UPDATE books SET available_copies = ? WHERE id = ?`, availableCopies-1, bookId)
		if err != nil {
			logger.Error("Error updating book availability: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update book availability"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Book borrowed successfully"})
	}
}
