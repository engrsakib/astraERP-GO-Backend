package faq

type CreateFaqRequest struct {
	Headline string `json:"headline" binding:"required"`
	Photo    string `json:"photo"`
}

type UpdateFaqRequest struct {
	Headline *string `json:"headline"`
	Photo    *string `json:"photo"`
}


type PaginationQuery struct {
 	Page int `form:"page" binding:"omitempty,min=1"`
 	Limit int `form:"limit" binding:"omitempty,min=1,max=100" default:"10"` 
	Search string `form:"search" binding:"omitempty"`
}
