package logdata

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logdata(c *gin.Context) {
	clientIP := c.ClientIP()
	fmt.Println("Client IP:", clientIP)
}
