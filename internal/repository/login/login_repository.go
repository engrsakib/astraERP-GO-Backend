package login

import (
    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
)

type LoginRepository struct {
    DB *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
    return &LoginRepository{DB: db}
}

func (r *LoginRepository) FindByMobile(mobile string) (*models.User, error) {
    var user models.User
    if err := r.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
