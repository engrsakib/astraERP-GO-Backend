package user

import (
	"errors"
	"os"
	"strconv"

	"github.com/engrsakib/erp-system/internal/models"
	"github.com/engrsakib/erp-system/internal/utils"
	userRepo "github.com/engrsakib/erp-system/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB   *gorm.DB
	Repo *userRepo.UserRepository
}

func NewUserService(db *gorm.DB, repo *userRepo.UserRepository) *UserService {
	return &UserService{DB: db, Repo: repo}
}


func (service *UserService) RegisterUser(authHeader, name, email, password, confirm string) (string, string, error) {

	payload, err := utils.ExtractTokenPayload(authHeader)
	if err != nil {
		return "", "", err
	}

	mobile, ok := payload["mobile"].(string)
	if !ok || mobile == "" {
		return "", "", errors.New("mobile not found in token")
	}

	
	if err := utils.ValidatePassword(password); err != nil {
		return "", "", err
	}

	if password != confirm {
		return "", "", errors.New("passwords do not match")
	}

	
	costStr := os.Getenv("BCRYPT_COST")
	cost, _ := strconv.Atoi(costStr)
	
	if cost < 10 {
		cost = bcrypt.DefaultCost
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", "", errors.New("failed to hash password")
	}

	hashed := string(hashedBytes)

	// ৪. ইউজার অবজেক্ট তৈরি
	user := models.User{
		Name:     name,
		Email:    email,
		Mobile:   mobile,
		Password: hashed,
	}

	// ৫. ডাটাবেসে ইউজার সেভ করা
	if err := service.DB.Create(&user).Error; err != nil {
		return "", "", err
	}

	// ৬. রেজিস্ট্রেশন সফল হলে টোকেন জেনারেট করা
	// নোট: user অবজেক্টে এখন ID চলে এসেছে (AutoIncrement এর কারণে)
	accessToken, err := utils.GenerateToken(payload, "ACCESS_TOKEN_EXP")
	if err != nil {
		return "", "", errors.New("failed to generate access token")
	}

		refreshToken, err := utils.GenerateToken(payload, "REFRESH_TOKEN_EXP")
	if err != nil {
		return "", "", errors.New("failed to generate refresh token")
	}

	// ৭. টোকেন রিটার্ন করা
	return accessToken, refreshToken, nil
}