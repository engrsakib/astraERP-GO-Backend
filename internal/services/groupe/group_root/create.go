package group_root

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type CreateGroupService struct {
    Repo *repo.GroupRepository
}

func NewCreateGroupService(r *repo.GroupRepository) *CreateGroupService {
    return &CreateGroupService{Repo: r}
}

func (s *CreateGroupService) Execute(req group.GroupRequest, addedBy int64) (*models.Group, error) {
    g := models.Group{
        Name:        req.Name,
        GroupTypeID: req.GroupTypeID,
        AddedBy:     addedBy,
    }

    if err := s.Repo.Create(&g); err != nil {
        return nil, err
    }

    return &g, nil
}
