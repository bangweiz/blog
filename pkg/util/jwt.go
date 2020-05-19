package util

import (
	"github.com/bangweiz/blog/pkg"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(pkg.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(username, password, email string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(6 * time.Hour)

	claims := Claims{
		username,
		password,
		email,
		jwt.StandardClaims {
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(tokenString string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
