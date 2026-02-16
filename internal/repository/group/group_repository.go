package group

import (
    "strings"

    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
)

type GroupRepository struct {
    db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
    return &GroupRepository{db: db}
}

func (r *GroupRepository) Create(group *models.Group) error {
    return r.db.Create(group).Error
}

func (r *GroupRepository) GetByID(id int64) (*models.Group, error) {
    var group models.Group
    err := r.db.First(&group, id).Error
    return &group, err
}

func (r *GroupRepository) Update(group *models.Group) error {
    return r.db.Save(group).Error
}

func (r *GroupRepository) Delete(id int64) error {
    result := r.db.Delete(&models.Group{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return gorm.ErrRecordNotFound
    }
    return nil
}

func (r *GroupRepository) GetAllWithQuery(search string, page, limit int) ([]models.Group, int64, error) {
    var items []models.Group
    var total int64

    query := r.db.Model(&models.Group{})

    if search != "" {
        search = strings.ToLower(search)
        query = query.Where("LOWER(name) LIKE ?", "%"+search+"%")
    }

    query.Count(&total)

    offset := (page - 1) * limit

    err := query.
        Order("updated_at DESC").
        Offset(offset).
        Limit(limit).
        Preload("GroupType").
        Preload("Creator").
        Find(&items).Error

    return items, total, err
}
