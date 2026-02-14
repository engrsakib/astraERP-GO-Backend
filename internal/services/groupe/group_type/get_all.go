package group_type

import (
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type GetAllGroupTypeService struct {
    Repo *repo.GroupTypeRepository
}

func NewGetAllGroupTypeService(r *repo.GroupTypeRepository) *GetAllGroupTypeService {
    return &GetAllGroupTypeService{Repo: r}
}

func (s *GetAllGroupTypeService) Execute() ([]models.GroupType, error) {
    return s.Repo.GetAll()
}
