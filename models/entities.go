package models

// User represents a user in the system
type User struct {
	ID           int    `json:"id" db:"id"`           // Unique identifier for the user
	Email        string `json:"email" db:"email"`     // User email address
	PasswordHash string `json:"-" db:"password_hash"` // Hashed password (not included in JSON responses)
	Name         string `json:"name" db:"name"`       // User's full name
}
