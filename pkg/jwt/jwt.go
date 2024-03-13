package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpDuration = time.Hour * 2

var jwtKey = []byte("constantine")

func GenToken(userID int64, username string) (string, error) {
	// Create the Claims
	c := Claims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "speakpeak",
			Subject:   "user token",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	// Get claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
