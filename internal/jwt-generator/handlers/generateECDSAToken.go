package handlers

import (
	"crypto/ecdsa"
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
// i havn't seen a memory corruptioninfo leakage in go yet
// but maybe it's a better idea, just to store the public ky in ram?
// and load the signKey on every signing request? depends on  your usge i guess
var (
	Ecdsa256PrivateSignKey *ecdsa.PrivateKey
	Ecdsa384PrivateSignKey *ecdsa.PrivateKey
	Ecdsa512PrivateSignKey *ecdsa.PrivateKey
)

// GenerateECDSAToken return an JWT token ECDSA encoded (ES256,ES...)
func GenerateECDSAToken(c *gin.Context) {

	// variable to store the encode method from the request
	// and optinal claims for map
	var data = new(dtos.DataToEncodeDto)

	guid := uuid.New().String()

	err := c.BindJSON(&data)
	if err != nil {
		// If the structure of the ody is wrong, return an HTTP error
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body, plese send a new request",
			"error":   err.Error(),
		})
		return
	}

	now := time.Now()

	// eclare the expiration time of the token
	// if otinal variable is present uses then else 5 minutes
	expirationTime := now.Add(5 * time.Minute)
	if data.ExpiresIn != nil {
		expirationTime = now.Add(time.Duration(*data.ExpiresIn) * time.Second)
	}

	// Create the JWT claims, which inclues the username and expiry time
	claims := jwt.StandardClaims{
		// In JWT, the expiry tie is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  now.Unix(),
		Id:        guid,
	}

	// map optional claims(data) to token claims(claims)
	err = copier.Copy(&claims, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error when mpping optional claims",
			"error":   err.Error(),
		})
		return
	}

	// Setting the nbf claim if its comes
	if data.NotBefore != nil {
		claims.NotBefore = now.Add(time.Duration(*data.NotBefore) * time.Second).Unix()
	}

	// Save the correct key for encode
	var key = new(ecdsa.PrivateKey)

	// Declare the oken with the algorithm used for signing, and the claims
	var token = new(jwt.Token)
	switch strings.ToUpper(data.Method) {
	case jwt.SigningMethodES256.Name:
		token = jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		key = Ecdsa256PrivateSignKey
	case jwt.SigningMethodES384.Name:
		token = jwt.NewWithClaims(jwt.SigningMethodES384, claims)
		key = Ecdsa384PrivateSignKey
	case jwt.SigningMethodES512.Name:
		token = jwt.NewWithClaims(jwt.SigningMethodES512, claims)
		key = Ecdsa512PrivateSignKey
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "please use a vali ECDSA encode method",
			"error":   errors.New("invalid encode method").Error(),
		})
		return
	}

	// Create the JWT string
	tokenString, err := token.SignedString(key)
	if err != nil {
		// If there is an error in creating the JWT return an intrnal server error
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating jwt",
			"error":   err.Error(),
		})
		return
	}

	// Finally, we set the client header "Authorizarion" as te JWT we just generated
	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, gin.H{
		"message": "token returned, see the Authorizarion header",
		"error":   nil,
	})
	return
}
