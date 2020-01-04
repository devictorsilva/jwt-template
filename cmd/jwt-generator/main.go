package main

import (
	"github.com/gin-gonic/gin"
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

	r.GET("/hmac-token", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hmac-token!",
			"error":   nil,
		})
	})
	r.GET("/rsa-token", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "rsa-token!",
			"error":   nil,
		})
	})
	r.GET("/ecdsa-token", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ecdsa-token!",
			"error":   nil,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
