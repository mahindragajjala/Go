package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsVerified   bool      `json:"is_verified"`
	Status       string    `json:"status"`
	LastLogin    time.Time `json:"last_login"`
	LoginCount   int       `json:"login_count"`
	Role         string    `json:"role"`
}
type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
