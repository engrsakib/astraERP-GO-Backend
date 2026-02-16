package group_root

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type UpdateGroupService struct {
    Repo *repo.GroupRepository
}

func NewUpdateGroupService(r *repo.GroupRepository) *UpdateGroupService {
    return &UpdateGroupService{Repo: r}
}

func (s *UpdateGroupService) Execute(id int64, req group.GroupRequest) error {
    g, err := s.Repo.GetByID(id)
    if err != nil {
        return err
    }

    g.Name = req.Name
    g.GroupTypeID = req.GroupTypeID

    return s.Repo.Update(g)
}
