package faq

import (
	dto "github.com/engrsakib/erp-system/internal/dto/faq"
	"github.com/engrsakib/erp-system/internal/models"
)


func (s *FaqService) CreateFaq(userID int64, req dto.CreateFaqRequest) (*dto.FaqResponse, error) {
	
	faq := models.Faq{
		Headline: req.Headline,
		Photo:    req.Photo,
		AddedBy:  &userID,
		// FaqAnswers 
	}

	if err := s.Repo.CreateFaq(&faq); err != nil {
		return nil, err
	}

	return &dto.FaqResponse{
		ID:        faq.ID,
		Headline:  faq.Headline,
		Photo:     faq.Photo,
		AddedBy:   faq.AddedBy,
		CreatedAt: faq.CreatedAt,
	}, nil
}