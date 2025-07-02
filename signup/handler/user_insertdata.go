package handler

import (
	"database/sql"
	"net/http"
	"time"

	SignupRequest "signup/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(c *gin.Context) {
	var req SignupRequest.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=auth_db sslmode=disable")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB connection failed"})
		return
	}
	defer db.Close()

	userID := uuid.New()
	now := time.Now()

	query := `
		INSERT INTO users (
			id, email, password_hash, created_at, updated_at, is_verified, status, last_login, login_count, role
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err = db.Exec(query,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB insert failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User signed up successfully"})
}
