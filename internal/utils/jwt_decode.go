package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractTokenPayload(authHeader string) (map[string]interface{}, error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("invalid authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	secret := os.Getenv("JWT_SECRET")

	// 🛠️ পরিবর্তন ১: Leeway যোগ করা হলো
	// ঘড়ির সময়ের একটু এদিক-সেদিক হলেও (৫ সেকেন্ড) এটামেনে নিবে।
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	}, jwt.WithLeeway(5*time.Second)) // ৫ সেকেন্ড ছাড় দেওয়া হলো

	// ডিবাগিং এর জন্য
	fmt.Println("DEBUG: Token String =", tokenStr)
	
	if err != nil {
		// 🛠️ পরিবর্তন ২: আসল এররটি দেখা খুবই জরুরি
		fmt.Println("❌ Token Parse Error:", err)
		return nil, fmt.Errorf("token parsing failed: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// 🛠️ পরিবর্তন ৩: টোকেনের মেয়াদ কনসোলে চেক করুন
	if exp, ok := claims["exp"].(float64); ok {
		fmt.Printf("ℹ️ Token Expires at: %v | Current Time: %v\n", time.Unix(int64(exp), 0), time.Now())
	}

	return claims, nil
}