package utils

import (
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

func GenerateTemporaryToken(mobile string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    expStr := os.Getenv("TEMPORARY_TOKEN_EXP")

    expMin, err := time.ParseDuration(expStr + "m")
    if err != nil {
        expMin = 15 * time.Minute
    }

    claims := jwt.MapClaims{
        "mobile": mobile,
        "exp":    time.Now().Add(expMin).Unix(),
        "type":   "temporary",
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
