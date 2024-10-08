package handler

import (
	"BookHaven/models"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// GetAllBooks retrieves a list of all books from the database.
// @Summary Get All Books
// @Description Retrieves details of all books including title, author, published year, ISBN, and available copies.
// @Tags Books
// @Produce json
// @Success 200 {array} models.BookDto
// @Failure 500 {object} models.ErrorResponse
// @Router /books [get]
func GetAllBooks(db *sql.DB, logger *logrus.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := db.Query("SELECT title, author, published_year, isbn, available_copies FROM books")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving books"})
		}
		defer rows.Close()

		var books []models.BookDto
		for rows.Next() {
			var book models.BookDto
			if err := rows.Scan(&book.Title, &book.Author, &book.PublishedYear, &book.ISBN, &book.AvailableCopies); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error scanning book"})
			}
			books = append(books, book)
		}

		if err := rows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error iterating books"})
		}

		return c.JSON(http.StatusOK, books)
	}
}
