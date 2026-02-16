package common

type PaginationQuery struct {
    Page  int `form:"page"`
    Limit int `form:"limit"`
    Search string `form:"search"`
}

type PaginationMeta struct {
    CurrentPage int `json:"current_page"`
    TotalPages  int `json:"total_pages"`
    TotalItems  int64 `json:"total_items"`
    Limit       int `json:"limit"`
}
