package faq

import (
	faqRepo "github.com/engrsakib/erp-system/internal/repository/faq"
)

type FaqService struct {
	Repo *faqRepo.FaqRepository
}

func NewFaqService(repo *faqRepo.FaqRepository) *FaqService {
	return &FaqService{Repo: repo}
}