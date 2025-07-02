// handlers/signup.go
package handler

import (
	"net/http"
	"time"

	"mongo/db"
	"mongo/models"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignupHandler(c *gin.Context) {
	var req SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Build audit log
	log := models.AuditLog{
		Email:     req.Email,
		IP:        c.ClientIP(),
		Action:    "signup",
		Timestamp: time.Now(),
	}

	if err := db.InsertAuditLog(log); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert audit log"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
