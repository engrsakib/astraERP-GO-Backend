package user

import (
	"errors"
	"os"
	"strconv"

	"github.com/engrsakib/erp-system/internal/models"
	"github.com/engrsakib/erp-system/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
    userRepo "github.com/engrsakib/erp-system/internal/repository/user"
)

type UserService struct {
    DB *gorm.DB
    Repo        *userRepo.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{DB: db}
}

func (service *UserService) RegisterUser(authHeader, name, email, password, confirm string) error {
    payload, err := utils.ExtractTokenPayload(authHeader)
    if err != nil {
        return err
    }

    mobile, ok := payload["mobile"].(string)
    if !ok || mobile == "" {
        return errors.New("mobile not found in token")
    }

    if err := utils.ValidatePassword(password); err != nil {
        return err
    }

    if password != confirm {
        return errors.New("passwords do not match")
    }

    // 🔐 Password Hashing with ENV-based cost
    costStr := os.Getenv("BCRYPT_COST")
    cost, _ := strconv.Atoi(costStr)
    if cost < 10 {
        cost = bcrypt.DefaultCost
    }

    hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
    if err != nil {
        return errors.New("failed to hash password")
    }

    hashed := string(hashedBytes)

    user := models.User{
        Name:     name,
        Email:    email,
        Mobile:   mobile,
        Password: hashed,
    }

    return service.DB.Create(&user).Error
}



func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}


// func checkPasswordHash(password, hash string) bool {
//     err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//     return err == nil
// }
