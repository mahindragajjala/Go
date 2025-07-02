// main.go
package main

import (
	"log"

	"mongo/db"
	"mongo/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init MongoDB
	if err := db.InitMongo(); err != nil {
		log.Fatalf("MongoDB init failed: %v", err)
	}

	r := gin.Default()

	r.POST("/signup", handler.SignupHandler)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
