package group_root

import (
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type GetGroupByIDService struct {
    Repo *repo.GroupRepository
}

func NewGetGroupByIDService(r *repo.GroupRepository) *GetGroupByIDService {
    return &GetGroupByIDService{Repo: r}
}

func (s *GetGroupByIDService) Execute(id int64) (*models.Group, error) {
    return s.Repo.GetByID(id)
}
