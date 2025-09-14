package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var jwtSecret = []byte(GetJWTSecret())

// Get Secret from Env (or fallback)
func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "defaultsecretkeyhere"
	}
	return secret
}

// Generate Token
func GenerateToken(userUUID string) (string, error) {
	claims := jwt.MapClaims{
		"uuid": userUUID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // 24 hr expiry
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
