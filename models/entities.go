package models

import "time"

type User struct {
	ID           int    // ID is the primary key, auto-incremented
	Email        string // User's email, must be unique
	PasswordHash string // Hashed password
	Name         string // User's name
	Balance      int    // User balance for borrowing books, default is 0
}

type Book struct {
	ID              int    // ID is the primary key, auto-incremented
	Title           string // Title of the book
	Author          string // Author of the book
	PublishedYear   int    // Year the book was published
	ISBN            string // ISBN of the book, must be unique
	AvailableCopies int    // Number of available copies, default is 1
}

type Borrowing struct {
	ID         int        // ID is the primary key, auto-incremented
	UserID     int        // Foreign key reference to users table
	BookID     int        // Foreign key reference to books table
	BorrowedAt time.Time  // Timestamp for when the book was borrowed
	ReturnedAt *time.Time // Nullable timestamp for when the book was returned
	Status     string     // Status of the borrowing (either 'borrowed' or 'returned')
}

type Transaction struct {
	ID            int    // ID is the primary key, auto-incremented
	UserID        int    // Foreign key reference to users table
	Amount        int    // The top-up amount
	TransactionID string // External transaction ID from payment provider, must be unique
	Status        string // Status of the top-up transaction (either 'pending', 'completed', or 'failed')
	Description   string
	CreatedAt     time.Time // Timestamp for when the top-up was created
	UpdatedAt     time.Time // Timestamp for when the top-up status was last updated
}
