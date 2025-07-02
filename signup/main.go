package main

import (
	"signup/db"
	routes "signup/handler"
	"signup/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectPostgres()

	r := gin.Default()

	// Apply middlewares globally
	r.Use(middlewares.Logger())
	r.Use(middlewares.CORSMiddleware())

	// Routes
	r.POST("/signup", routes.Signup)

	r.Run(":8080")
}
