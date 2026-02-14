package faq

import (
	"github.com/engrsakib/erp-system/internal/models"
	"gorm.io/gorm"
)

type FaqRepository struct {
	DB *gorm.DB
}

func NewFaqRepository(db *gorm.DB) *FaqRepository {
	return &FaqRepository{DB: db}
}


func (r *FaqRepository) CreateFaq(faq *models.Faq) error {
	return r.DB.Create(faq).Error
}

func (r *FaqRepository) CreateAnswer(answer *models.FaqAnswer) error {
	return r.DB.Create(answer).Error
}

func (r *FaqRepository) FindByID(id string) (*models.Faq, error) {
	var faq models.Faq
	err := r.DB.Preload("FaqAnswers").First(&faq, "id = ?", id).Error
	return &faq, err
}

func (r *FaqRepository) FindAll(page, limit int, search string) ([]models.Faq, int64, error) {
	var faqs []models.Faq
	var total int64
	query := r.DB.Model(&models.Faq{})

	if search != "" {
		query = query.Where("headline ILIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&faqs).Error
	
	return faqs, total, err
}


// func (r *FaqRepository) FindByID(id string) (*models.Faq, error) {
// 	var faq models.Faq
	
// 	err := r.DB.Preload("FaqAnswers").First(&faq, "id = ?", id).Error
// 	return &faq, err
// }

// Update
func (r *FaqRepository) Update(faq *models.Faq) error {
	return r.DB.Save(faq).Error
}


func (r *FaqRepository) Delete(id string) error {
	return r.DB.Delete(&models.Faq{}, "id = ?", id).Error
}