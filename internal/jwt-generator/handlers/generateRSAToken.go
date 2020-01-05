package handlers

import (
	"github.com/gin-gonic/gin"
)

// GenerateRSAToken return an JWT token RSA encoded (RS256,RS...)
func GenerateRSAToken(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "rsa-token!",
		"error":   nil,
	})
}
