package group_root

import (
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type DeleteGroupService struct {
    Repo *repo.GroupRepository
}

func NewDeleteGroupService(r *repo.GroupRepository) *DeleteGroupService {
    return &DeleteGroupService{Repo: r}
}

func (s *DeleteGroupService) Execute(id int64) error {
    return s.Repo.Delete(id)
}
