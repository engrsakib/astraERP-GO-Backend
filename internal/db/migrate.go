package db

import (
	"log"

	"gorm.io/gorm"
	"github.com/engrsakib/erp-system/internal/models"
)


func Migrate(db *gorm.DB) error {
	log.Println("🔄 Running Database Migration...")

	
	err := db.AutoMigrate(
		&models.User{},           
		&models.UserPermission{}, 
		// ভবিষ্যতে: &models.Product{},
		// ভবিষ্যতে: &models.Order{},
	)

	if err != nil {
		log.Printf("⚠️ Warning: AutoMigrate failed: %v", err)
		return err
	}

	log.Println("✅ Database Migration Successful")
	return nil
}