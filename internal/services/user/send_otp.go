package user

import (
    "context"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"

    "github.com/redis/go-redis/v9"
    "github.com/engrsakib/erp-system/internal/utils"
)

type OTPService struct {
    Redis *redis.Client
}

func NewOTPService(rdb *redis.Client) *OTPService {
    return &OTPService{Redis: rdb}
}

func (s *OTPService) GenerateOTP() string {
    rand.Seed(time.Now().UnixNano())
    return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func (service *OTPService) SendOTP(mobile string) error {
    otp := service.GenerateOTP()
    key := "OTP:" + mobile

    // TTL from env
    expStr := os.Getenv("OTP_TOKEN_EXP")
    expMin, err := strconv.Atoi(expStr)
    if err != nil || expMin <= 0 {
        expMin = 2
    }

    
    err = service.Redis.Set(context.Background(), key, otp, time.Duration(expMin)*time.Minute).Err()
    if err != nil {
        return fmt.Errorf("failed to store OTP: %v", err)
    }

    // realValueInRedis, _ := service.Redis.Get(context.Background(), key).Result()
    // // fmt.Printf("✅ SUCCESS: Redis এ ডাটা আছে! Key: %s | Real Value: %s\n", key, realValueInRedis)

    // Send SMS
    fmt.Println("otp sent:", otp)
    return utils.SendSMS(otp, []string{mobile})
}
