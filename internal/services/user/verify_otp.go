package user

import (
    "context"
    "fmt"

    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/redis/go-redis/v9"
)

func (service *OTPService) VerifyOTP(mobile string, otp string) (string, error) {
	key := "OTP:" + mobile
    // fmt.Printf("DEBUG INPUT: Mobile='%s' | Key='%s' | OTP='%s'\n", mobile, key, otp)

    storedOTP, err := service.Redis.Get(context.Background(), key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("otp expired or not found")
    }
    if err != nil {
        return "", fmt.Errorf("failed to read otp")
    }

	// fmt.Printf("DEBUG CHECK: Stored='%s' | Input='%s'\n", storedOTP, otp)

    if storedOTP != otp {
        return "", fmt.Errorf("invalid otp")
    }

    service.Redis.Del(context.Background(), key)

    token, err := utils.GenerateTemporaryToken(mobile)
    if err != nil {
        return "", fmt.Errorf("failed to generate token")
    }

    return token, nil
}
