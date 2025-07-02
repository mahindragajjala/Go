package main

import (
	"jwt/create"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Routes
	r.POST("/signup", create.LoginHandler)

	r.Run(":8080")
}
