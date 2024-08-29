package models

import "time"

// User represents a user in the system
type RegisterDto struct {
	Email    string `json:"email"`    // User email address
	Password string `json:"password"` // User Password
	Name     string `json:"name" `    // User's full name
}

type LoginDto struct {
	Email    string `json:"email"` // User email address
	Password string `json:"password"`
}

type TopUpRequestDto struct {
	Amount uint `json:"amount"`
}

type ErrorResponseDto struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type BookDto struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublishedYear   int    `json:"published_year"`
	ISBN            string `json:"isbn"`
	AvailableCopies int    `json:"available_copies"`
}

type BorrowBookDto struct {
	ISBN string `json:"isbn"`
}

type ReturnBookDto struct {
	ISBN string `json:"isbn"`
}

type BorrowedBookResponse struct {
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	ISBN       string    `json:"isbn"`
	BorrowedAt time.Time `json:"borrowed_at"`
}

type JokeDto struct {
	Joke string `json:"joke"`
}
