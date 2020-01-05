package middleware

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	HmacPublicSignKey      = []byte("my_secret_key")
	RsaPrivateSignKey      *rsa.PrivateKey
	Ecdsa256PrivateSignKey *ecdsa.PrivateKey
	Ecdsa384PrivateSignKey *ecdsa.PrivateKey
	Ecdsa512PrivateSignKey *ecdsa.PrivateKey
)

// Auth validate the token string
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		hasPrefix := strings.HasPrefix(token, "Bearer ")
		if !hasPrefix {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "the Authorization Header MUST BE Authorization: Bearer <T>",
				"error":   errors.New("invalid Authorization schema").Error(),
			})
			return
		}
		token = strings.ReplaceAll(token, "Bearer ", "")

		var claims = new(jwt.StandardClaims)

		fmt.Print(token)
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {

			switch token.Method.Alg() {
			case jwt.SigningMethodES256.Name:
				return Ecdsa256PrivateSignKey.PublicKey, nil
			case jwt.SigningMethodES384.Name:
				return Ecdsa384PrivateSignKey.PublicKey, nil
			case jwt.SigningMethodES512.Name:
				return Ecdsa512PrivateSignKey.PublicKey, nil
			case jwt.SigningMethodHS256.Name:
				return HmacPublicSignKey, nil
			case jwt.SigningMethodHS384.Name:
				return HmacPublicSignKey, nil
			case jwt.SigningMethodHS512.Name:
				return HmacPublicSignKey, nil
			case jwt.SigningMethodRS256.Name:
				return RsaPrivateSignKey.PublicKey, nil
			case jwt.SigningMethodRS384.Name:
				return RsaPrivateSignKey.PublicKey, nil
			case jwt.SigningMethodRS512.Name:
				return RsaPrivateSignKey.PublicKey, nil
			default:
				err := errors.New("invalid encode method")
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"message": "please use a valid encode method",
					"error":   err.Error(),
				})
				return nil, err
			}
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "please use a valid encode method",
					"error":   err.Error(),
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "is that an token?",
				"error":   err.Error(),
			})
			return
		}
		if !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token is not valid",
				"error":   errors.New("invalid token").Error(),
			})
			return
		}
		return
	}
}
