package main

import (
	"github.com/gin-gonic/gin"

	"jwt-template/internal/jwt-generator/handlers"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
			"error":   nil,
		})
	})

	r.GET("/hmac-token", handlers.GenerateHMACToken)
	r.GET("/rsa-token", handlers.GenerateRSAToken)
	r.GET("/ecdsa-token", handlers.GenerateECDSAToken)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
