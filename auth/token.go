package auth

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Custom struct {
	AddTime time.Duration
}

type CustomClaims struct {
	Custom
	jwt.RegisteredClaims
}

// 签名
func (c *CustomClaims) GenToken(signingKey string) (string, error) {
	// Create the claims
	claims := CustomClaims{
		c.Custom,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(c.AddTime * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    c.Issuer,
			Subject:   c.Subject,
			ID:        c.ID,
			Audience:  c.Audience,
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(signingKey)
}

// 验证签名
func AuthToken(tokenString string, signingKey string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if token.Valid {
		fmt.Println("You look nice today")
		return nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		fmt.Println("Timing is everything")
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
	return err
}

func Done(w http.ResponseWriter, r *http.Request) bool {
	return true
}
