package faq

import (
	dto "github.com/engrsakib/erp-system/internal/dto/faq"
)


func (s *FaqService) UpdateFaq(id string, req dto.UpdateFaqRequest) (*dto.FaqResponse, error) {
    
    faq, err := s.Repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    
    if req.Headline != nil {
        faq.Headline = *req.Headline
    }
    if req.Photo != nil {
        faq.Photo = *req.Photo
    }

    if err := s.Repo.Update(faq); err != nil {
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
