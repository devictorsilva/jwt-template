package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"jwt-template/internal/jwt-validator/middleware"
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

		middleware.RsaPrivateSignKey, err = jwt.ParseRSAPrivateKeyFromPEM(rsaPrivateSignBytes)
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

		middleware.Ecdsa256PrivateSignKey, err = jwt.ParseECPrivateKeyFromPEM(ecdsa256PrivateSignBytes)
		if err != nil {
			fmt.Printf("Error when set ECDSA256 Private Key: %+v", err)
			os.Exit(1)
		}

		ecdsa384PrivateSignBytes, err := ioutil.ReadFile(ecdsa384PrivateKeyPath)
		if err != nil {
			fmt.Printf("Error when open ECDSA384 Private Key: %+v", err)
			os.Exit(1)
		}

		middleware.Ecdsa384PrivateSignKey, err = jwt.ParseECPrivateKeyFromPEM(ecdsa384PrivateSignBytes)
		if err != nil {
			fmt.Printf("Error when set ECDSA384 Private Key: %+v", err)
			os.Exit(1)
		}

		ecdsa512PrivateSignBytes, err := ioutil.ReadFile(ecdsa512PrivateKeyPath)
		if err != nil {
			fmt.Printf("Error when open ECDSA512 Private Key: %+v", err)
			os.Exit(1)
		}

		middleware.Ecdsa512PrivateSignKey, err = jwt.ParseECPrivateKeyFromPEM(ecdsa512PrivateSignBytes)
		if err != nil {
			fmt.Printf("Error when set ECDSA512 Private Key: %+v", err)
			os.Exit(1)
		}
	}
	// (ECDSA) END
}

func main() {
	r := gin.Default()

	logger, _ := zap.NewProduction()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
			"error":   nil,
		})
	})

	auth := r.Group("/private")
	auth.Use(middleware.Auth())

	auth.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "authenticated pong!",
			"error":   nil,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
