package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

// Claims ...
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// ParseToken ...
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claim, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claim, nil
		}
	}
	return nil, err
}

// GenerateToken ...
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * time.Minute)

	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),

		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-learning",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
