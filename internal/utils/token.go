package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload map[string]interface{}, envKey string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    expStr := os.Getenv(envKey)

    if secret == "" || expStr == "" {
        return "", jwt.ErrTokenMalformed
    }

    expMin, err := strconv.Atoi(expStr)
    if err != nil || expMin <= 0 {
        return "", jwt.ErrTokenExpired
    }

    claims := jwt.MapClaims{}
    for k, v := range payload {
        claims[k] = v
    }
    claims["exp"] = time.Now().Add(time.Duration(expMin) * time.Minute).Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
