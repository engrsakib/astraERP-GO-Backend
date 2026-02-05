package login

import (
    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
	"fmt"
)

type LoginRepository struct {
    DB *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
    return &LoginRepository{DB: db}
}

// func (r *LoginRepository) FindByMobile(mobile string) (*models.User, error) {
//     var user models.User
//     if err := r.DB.Where("mobile = ?", mobile).First(&user).Error; err != nil {
//         return nil, err
//     }
	
//     return &user, nil
// }


func (r *LoginRepository) FindByMobile(mobile string) (*models.User, error) {
    var user models.User

    // fmt.Println("========== LOGIN REPOSITORY DEBUG ==========")
    // fmt.Printf("➡️  Query Mobile: '%s'\n", mobile)
    // fmt.Printf("➡️  Length: %d\n", len(mobile))

    result := r.DB.Where("mobile = ?", mobile).First(&user)

    if result.Error != nil {
        fmt.Printf("❌ DB ERROR: %v\n", result.Error)
        fmt.Println("============================================")
        return nil, result.Error
    }

    // fmt.Println("✅ User Found:")
    // fmt.Printf("   ID: %d\n", user.ID)
    // fmt.Printf("   Name: %s\n", user.Name)
    // fmt.Printf("   Mobile: %s\n", user.Mobile)
    // fmt.Printf("   Email: %s\n", user.Email)
    // fmt.Println("============================================")

    return &user, nil
}
