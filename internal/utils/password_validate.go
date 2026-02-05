package utils

import (
    "errors"
    "regexp"
)

func ValidatePassword(password string) error {
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters")
    }
    if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
        return errors.New("password must contain at least one uppercase letter")
    }
    if !regexp.MustCompile(`[0-9]`).MatchString(password) {
        return errors.New("password must contain at least one digit")
    }
    if !regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=

\[\]

\{\};:'",<>\.\?/\\|]`).MatchString(password) {
        return errors.New("password must contain at least one special character")
    }
    return nil
}
