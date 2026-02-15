package group_root

import (
    "math"

    "github.com/engrsakib/erp-system/internal/dto/common"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type GetAllGroupService struct {
    Repo *repo.GroupRepository
}

func NewGetAllGroupService(r *repo.GroupRepository) *GetAllGroupService {
    return &GetAllGroupService{Repo: r}
}

func (s *GetAllGroupService) Execute(q common.PaginationQuery) ([]models.Group, common.PaginationMeta, error) {

    if q.Page == 0 {
        q.Page = 1
    }
    if q.Limit == 0 {
        q.Limit = 10
    }

    items, total, err := s.Repo.GetAllWithQuery(q.Search, q.Page, q.Limit)
    if err != nil {
        return nil, common.PaginationMeta{}, err
    }

    meta := common.PaginationMeta{
        CurrentPage: q.Page,
        Limit:       q.Limit,
        TotalItems:  total,
        TotalPages:  int(math.Ceil(float64(total) / float64(q.Limit))),
    }

    return items, meta, nil
}
