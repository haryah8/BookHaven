package models

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
