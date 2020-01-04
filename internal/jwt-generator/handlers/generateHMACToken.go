package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jwt-template/internal/jwt-generator/models/dtos"
)

// HmacPrivateSignKey is the private key in memory
var HmacPrivateSignKey = []byte("my_secret_key")

// GenerateHMACToken return an JWT token HMAC encoded (HS256,HS...)
func GenerateHMACToken(c *gin.Context) {

	// variable to store the encode method from the request
	// and optinal claims for map
	var data = new(dtos.DataToEncodeDto)

	err := c.BindJSON(&data)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body, plese send a new request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": data,
		"error":   nil,
	})
}
