package permission

import (
	"github.com/engrsakib/erp-system/internal/models"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	DB *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}


func (r *PermissionRepository) SyncPermissions(userID int64, slugs []string) error {
	
	return r.DB.Transaction(func(tx *gorm.DB) error {
		
		
		// DELETE FROM user_permissions WHERE user_id = ?
		if err := tx.Where("user_id = ?", userID).Delete(&models.UserPermission{}).Error; err != nil {
			return err
		}

		
		if len(slugs) == 0 {
			return nil
		}

		
		var newPermissions []models.UserPermission
		for _, slug := range slugs {
			newPermissions = append(newPermissions, models.UserPermission{
				UserID:         userID,
				PermissionSlug: slug,
			})
		}

		
		if err := tx.Create(&newPermissions).Error; err != nil {
			return err
		}

		return nil 
	})
}