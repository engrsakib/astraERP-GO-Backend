package group_type

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type CreateGroupTypeService struct {
    Repo *repo.GroupTypeRepository
}

func NewCreateGroupTypeService(r *repo.GroupTypeRepository) *CreateGroupTypeService {
    return &CreateGroupTypeService{Repo: r}
}

func (s *CreateGroupTypeService) Execute(req group.GroupTypeRequest) (*models.GroupType, error) {
    gt := models.GroupType{Name: req.Name}
    if err := s.Repo.Create(&gt); err != nil {
        return nil, err
    }
    return &gt, nil
}
