package handlers

import (
	"github.com/gin-gonic/gin"
)

// GenerateECDSAToken return an JWT token ECDSA encoded (ES256,ES...)
func GenerateECDSAToken(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "ecdsa-token!",
		"error":   nil,
	})
}
