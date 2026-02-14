package faq

import "time"

type FaqAnswerResponse struct {
	ID       int64  `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type FaqLinks struct {
    Self   string `json:"self"`
    Edit   string `json:"edit"`
    Delete string `json:"delete"`
}

type FaqResponse struct {
	ID        int64     `json:"id"`
	Headline string    `json:"headline"`
	Photo     string    `json:"photo"`
	AddedBy   *int64    `json:"added_by"`
	Answers  []FaqAnswerResponse `json:"answers,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateAnswerRequest struct {
	FaqID    int64  `json:"faq_id" binding:"required"` 
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
	Limit int `json:"limit"`
}
