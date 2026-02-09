package user

import "time"


type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	UserType  int8      `json:"user_type"`
	Photo     string    `json:"photo"`
	Permissions []string `json:"permissions,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}


type PaginationQuery struct {
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Search string `form:"search"` 
}


type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	TotalItems  int64 `json:"total_items"`
	Limit       int   `json:"limit"`
}