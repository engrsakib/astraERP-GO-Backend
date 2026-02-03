package user

import (
    "errors"

    "github.com/engrsakib/erp-system/internal/models"
    "github.com/engrsakib/erp-system/internal/utils"
    "gorm.io/gorm"
)

type UserService struct {
    DB *gorm.DB
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

    user := models.User{
        Name:     name,
        Email:    email,
        Mobile:   mobile,
        Password: password, // hashing later
    }

    return service.DB.Create(&user).Error
}
