package faq

import (
	"math"
	dto "github.com/engrsakib/erp-system/internal/dto/faq"
)


func (s *FaqService) GetFaqs(query dto.PaginationQuery) ([]dto.FaqResponse, *dto.PaginationMeta, error) {
	
	faqs, total, err := s.Repo.FindAll(query.Page, query.Limit, query.Search)
	if err != nil {
		return nil, nil, err
	}


	var response []dto.FaqResponse
	for _, f := range faqs {
		response = append(response, dto.FaqResponse{
			ID:        f.ID,
			Headline:  f.Headline,
			Photo:     f.Photo,
			AddedBy:   f.AddedBy,
			CreatedAt: f.CreatedAt,
			
		})
	}

	
	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
	meta := &dto.PaginationMeta{
		CurrentPage: query.Page,
		TotalPages:  totalPages,
		TotalItems:  total,
		Limit:       query.Limit,
	}

	return response, meta, nil
}


func (s *FaqService) GetFaqByID(id string) (*dto.FaqResponse, error) {
	
	faq, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	
	var answerResponses []dto.FaqAnswerResponse
	for _, ans := range faq.FaqAnswers {
		answerResponses = append(answerResponses, dto.FaqAnswerResponse{
			ID:       ans.ID,
			Question: ans.Question,
			Answer:   ans.Answer,
		})
	}

	
	return &dto.FaqResponse{
		ID:       faq.ID,
		Headline: faq.Headline,
		Photo:    faq.Photo,
		AddedBy:  faq.AddedBy,
		Answers:  answerResponses, 
		CreatedAt: faq.CreatedAt,
	}, nil
}