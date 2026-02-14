package faq

import (
	dto "github.com/engrsakib/erp-system/internal/dto/faq"
	"github.com/engrsakib/erp-system/internal/models"
)


func (s *FaqService) CreateAnswer(userID int64, req dto.CreateAnswerRequest) (*dto.FaqAnswerResponse, error) {
	
	
	answer := models.FaqAnswer{
		FaqID:    req.FaqID,    
		Question: req.Question,
		Answer:   req.Answer,
		AddedBy:  &userID,
	}

	
	if err := s.Repo.CreateAnswer(&answer); err != nil {
		return nil, err
	}

	
	return &dto.FaqAnswerResponse{
		ID:       answer.ID,
		Question: answer.Question,
		Answer:   answer.Answer,
	}, nil
}