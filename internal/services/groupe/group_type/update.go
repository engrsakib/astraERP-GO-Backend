package group_type

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type UpdateGroupTypeService struct {
    Repo *repo.GroupTypeRepository
}

func NewUpdateGroupTypeService(r *repo.GroupTypeRepository) *UpdateGroupTypeService {
    return &UpdateGroupTypeService{Repo: r}
}

func (s *UpdateGroupTypeService) Execute(id int64, req group.GroupTypeRequest) (*models.GroupType, error) {
    gt, err := s.Repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    gt.Name = req.Name

    if err := s.Repo.Update(gt); err != nil {
        return nil, err
    }

    return gt, nil
}
