package user

import (
	"github.com/engrsakib/erp-system/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}


func (r *UserRepository) GetUsers(page int, limit int, search string) ([]models.User, int64, error) {
	var users []models.User
	var total int64


	query := r.DB.Model(&models.User{})

	
	if search != "" {
		searchParam := "%" + search + "%"
		
		query = query.Where("name ILIKE ? OR mobile LIKE ? OR email ILIKE ?", searchParam, searchParam, searchParam)
	}

	
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("id DESC").Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// single user with permissions
func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	
	// ✅ use Preload to load permissions along with the user
	// here we assume that User model has a field named "Permissions" which is a slice of UserPermission
	// and that the UserPermission model has a foreign key relationship with User
	if err := r.DB.Preload("Permissions").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	
	return &user, nil
}