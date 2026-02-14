package group

import (
	"github.com/engrsakib/erp-system/internal/models"
	"gorm.io/gorm"
)

type GroupTypeRepository struct {
	db *gorm.DB
}

func NewGroupTypeRepository(db *gorm.DB) *GroupTypeRepository {
	return &GroupTypeRepository{db: db}
}

func (r *GroupTypeRepository) Create(groupType *models.GroupType) error {
	return r.db.Create(groupType).Error
}

func (r *GroupTypeRepository) GetAll() ([]models.GroupType, error) {
	var groupTypes []models.GroupType
	err := r.db.Find(&groupTypes).Error
	return groupTypes, err
}

func (r *GroupTypeRepository) GetByID(id int64) (*models.GroupType, error) {
	var groupType models.GroupType
	err := r.db.First(&groupType, id).Error
	return &groupType, err
}

func (r *GroupTypeRepository) Update(groupType *models.GroupType) error {
	return r.db.Save(groupType).Error
}

func (r *GroupTypeRepository) Delete(id int64) error {
	
	result := r.db.Delete(&models.GroupType{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}