package handlers

import (
	"crypto/rsa"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"jwt-template/internal/jwt-generator/models/dtos"
)

// keys are held in global variables
// i havn't seen a memory corruption/info leakage in go yet
// but maybe it's a better idea, just to store the public key in ram?
// and load the signKey on every signing request? depends on  your usage i guess
var (
	// VerifyKey *rsa.PublicKey
	RsaPrivateSignKey *rsa.PrivateKey
)

// GenerateRSAToken return an JWT token RSA encoded (RS256,RS...)
func GenerateRSAToken(c *gin.Context) {

	// variable to store the encode method from the request
	// and optinal claims for map
	var data = new(dtos.DataToEncodeDto)

	guid := uuid.New().String()

	err := c.BindJSON(&data)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body, plese send a new request",
			"error":   err.Error(),
		})
		return
	}

	now := time.Now()

	// Declare the expiration time of the token
	// if optinal variable is present uses then else 5 minutes
	expirationTime := now.Add(5 * time.Minute)
	if data.ExpiresIn != nil {
		expirationTime = now.Add(time.Duration(*data.ExpiresIn) * time.Second)
	}

	// Create the JWT claims, which includes the username and expiry time
	claims := jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  now.Unix(),
		Id:        guid,
	}

	// map optional claims(data) to token claims(claims)
	err = copier.Copy(&claims, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error when mapping optional claims",
			"error":   err.Error(),
		})
		return
	}

	// Setting the nbf claim if its comes
	if data.NotBefore != nil {
		claims.NotBefore = now.Add(time.Duration(*data.NotBefore) * time.Second).Unix()
	}

	// Declare the token with the algorithm used for signing, and the claims
	var token = new(jwt.Token)
	switch strings.ToUpper(data.Method) {
	case jwt.SigningMethodRS256.Name:
		token = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	case jwt.SigningMethodRS384.Name:
		token = jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
	case jwt.SigningMethodRS512.Name:
		token = jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please use a valid RSA encode method",
			"error":   errors.New("invalid encode method").Error(),
		})
		return
	}

	// Create the JWT string
	tokenString, err := token.SignedString(RsaPrivateSignKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating jwt",
			"error":   err.Error(),
		})
		return
	}

	// Finally, we set the client header "Authorizarion" as the JWT we just generated
	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{
		"message": "token returned, see the Authorizarion header",
		"error":   nil,
	})
	return
}
