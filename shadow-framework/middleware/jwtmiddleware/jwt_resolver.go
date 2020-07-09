package jwtmiddleware

import (
	"errors"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// JwtParser is a gin middleware for handle jwt token
func JwtParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtHeader := c.GetHeader("X-Custom-Jwt")
		Log.WithField("jwtHeader", jwtHeader).Debug()

		if strings.TrimSpace(jwtHeader) == "" {
			Log.Error("jwt token header not found")
			c.Error(errors.New("invalid token"))
		}

		error := parseToken(jwtHeader)
		if error != nil {
			c.Error(error)
		}

		c.Next()

	}
}

func parseToken(tokenString string) error {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	Log.WithField("tokenString", tokenString).Debug()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")

		hmacSampleSecret := []byte("mySecret")
		return hmacSampleSecret, nil
	})

	if err != nil {
		Log.Error(err)
		return err
	}

	Log.WithField("token", token).Debug()

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Log.WithField("claims", claims).Debug()
	} else {
		return errors.New("invalid token")
	}
	return nil
}

func createToken(secret []byte, claims jwt.MapClaims) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	Log.WithField("tokenString", tokenString).Debug("create a new jwt token")
	return tokenString
}
