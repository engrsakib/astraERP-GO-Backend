package faq

import "time"

type FaqAnswerResponse struct {
	ID       int64  `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}


type FaqResponse struct {
	ID        int64     `json:"id"`
	Headline string    `json:"headline"`
	Photo     string    `json:"photo"`
	AddedBy   *int64    `json:"added_by"`
	Answers  []FaqAnswerResponse `json:"answers,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}


type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalItems  int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
	Limit int `json:"limit"`
}
