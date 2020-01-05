package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"jwt-template/internal/jwt-generator/handlers"
)

// location of the files used for signing and verification
const (
	rsaPrivateKeyPath = "../../configs/rsa/private.rsa"

)

// read the key files before starting http handlers
func init() {
	// (RSA) BEGIN
	{
		rsaPrivateSignBytes, err := ioutil.ReadFile(rsaPrivateKeyPath)
		if err != nil {
			fmt.Printf("Error when open RSA Private Key: %+v", err)
			os.Exit(1)
		}

		handlers.RsaPrivateSignKey, err = jwt.ParseRSAPrivateKeyFromPEM(rsaPrivateSignBytes)
		if err != nil {
			fmt.Printf("Error when set RSA Private Key: %+v", err)
			os.Exit(1)
		}
	}
	// (RSA) END

}

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
