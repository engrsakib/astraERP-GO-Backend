package utils

import (
    "errors"
    "os"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

func ExtractTokenPayload(authHeader string) (map[string]interface{}, error) {
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        return nil, errors.New("invalid authorization header")
    }

    tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
    secret := os.Getenv("JWT_SECRET")

    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secret), nil
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid or expired token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, errors.New("invalid token claims")
    }

    return claims, nil
}
