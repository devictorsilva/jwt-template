package handlers

import (
	"github.com/gin-gonic/gin"
)

// GenerateHMACToken return an JWT token HMAC encoded (HS256,HS...)
func GenerateHMACToken(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "hmac-token!",
		"error":   nil,
	})
}
