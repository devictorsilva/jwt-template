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

	ecdsa256PrivateKeyPath = "../../configs/ed/ec256-private.pem"
	ecdsa384PrivateKeyPath = "../../configs/ed/ec384-private.pem"
	ecdsa512PrivateKeyPath = "../../configs/ed/ec512-private.pem"
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

	// (ECDSA) BEGIN
	{
		ecdsa256PrivateSignBytes, err := ioutil.ReadFile(ecdsa256PrivateKeyPath)
		if err != nil {
			fmt.Printf("Error when open ECDSA256 Private Key: %+v", err)
			os.Exit(1)
		}

		handlers.Ecdsa256PrivateSignKey, err = jwt.ParseECPrivateKeyFromPEM(ecdsa256PrivateSignBytes)
		if err != nil {
			fmt.Printf("Error when set ECDSA256 Private Key: %+v", err)
			os.Exit(1)
		}

		ecdsa384PrivateSignBytes, err := ioutil.ReadFile(ecdsa384PrivateKeyPath)
		if err != nil {
			fmt.Printf("Error when open ECDSA384 Private Key: %+v", err)
			os.Exit(1)
		}

		handlers.Ecdsa384PrivateSignKey, err = jwt.ParseECPrivateKeyFromPEM(ecdsa384PrivateSignBytes)
		if err != nil {
			fmt.Printf("Error when set ECDSA384 Private Key: %+v", err)
			os.Exit(1)
		}

		ecdsa512PrivateSignBytes, err := ioutil.ReadFile(ecdsa512PrivateKeyPath)
		if err != nil {
			fmt.Printf("Error when open ECDSA512 Private Key: %+v", err)
			os.Exit(1)
		}

		handlers.Ecdsa512PrivateSignKey, err = jwt.ParseECPrivateKeyFromPEM(ecdsa512PrivateSignBytes)
		if err != nil {
			fmt.Printf("Error when set ECDSA512 Private Key: %+v", err)
			os.Exit(1)
		}
	}
	// (ECDSA) END
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
