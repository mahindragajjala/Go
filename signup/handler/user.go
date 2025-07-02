package handler

import (
	"net/http"
	"signup/db"
	"signup/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var req models.SignupRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if user already exists
	var exists int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email=$1", req.Email).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists. Please login."})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	// Prepare data
	userID := uuid.New()
	now := time.Now()

	query := `
		INSERT INTO users (
			id, email, password_hash, created_at, updated_at, is_verified, status, last_login, login_count, role
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.DB.Exec(query,
		userID,
		req.Email,
		string(hashedPassword),
		now,
		now,
		false,       // is_verified
		"active",    // status
		time.Time{}, // last_login
		0,           // login_count
		"user",      // role
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}
