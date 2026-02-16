package group

import (
    "github.com/engrsakib/erp-system/internal/models"
    "gorm.io/gorm"
)

type MemberRepository struct {
    DB *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
    return &MemberRepository{DB: db}
}


func (r *MemberRepository) Create(member *models.Member) error {
    return r.DB.Create(member).Error
}


func (r *MemberRepository) GetAll(page, limit int, search string) ([]models.Member, int64, error) {
    var members []models.Member
    var total int64

    query := r.DB.Model(&models.Member{}).
        Preload("User").
        Preload("Group").
        Preload("Creator")

    if search != "" {
        query = query.Joins("JOIN users ON users.id = members.user_id").
            Where("users.name LIKE ?", "%"+search+"%")
    }

    query.Count(&total)

    offset := (page - 1) * limit
    err := query.Offset(offset).Limit(limit).Find(&members).Error

    return members, total, err
}


func (r *MemberRepository) GetByID(id int64) (*models.Member, error) {
    var member models.Member
    err := r.DB.Preload("User").
        Preload("Group").
        Preload("Creator").
        First(&member, id).Error
    return &member, err
}


func (r *MemberRepository) Update(member *models.Member) error {
    return r.DB.Save(member).Error
}


func (r *MemberRepository) Delete(id int64) error {
    return r.DB.Delete(&models.Member{}, id).Error
}
