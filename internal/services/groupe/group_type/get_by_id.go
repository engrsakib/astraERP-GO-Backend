package group_type

import (
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type GetGroupTypeByIDService struct {
    Repo *repo.GroupTypeRepository
}

func NewGetGroupTypeByIDService(r *repo.GroupTypeRepository) *GetGroupTypeByIDService {
    return &GetGroupTypeByIDService{Repo: r}
}

func (s *GetGroupTypeByIDService) Execute(id int64) (*models.GroupType, error) {
    return s.Repo.GetByID(id)
}
