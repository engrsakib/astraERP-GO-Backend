package group

import "github.com/engrsakib/erp-system/internal/dto/common"

type MemberRequest struct {
    UserID  int64 `json:"user_id"`
    GroupID int64 `json:"group_id"`
    Status  int8  `json:"status"`
}

type MemberSearchQuery struct {
    common.PaginationQuery
    Search string `form:"search"`
}
