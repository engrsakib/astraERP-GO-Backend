package group_type

import repo "github.com/engrsakib/erp-system/internal/repository/group"

type DeleteGroupTypeService struct {
    Repo *repo.GroupTypeRepository
}

func NewDeleteGroupTypeService(r *repo.GroupTypeRepository) *DeleteGroupTypeService {
    return &DeleteGroupTypeService{Repo: r}
}

func (s *DeleteGroupTypeService) Execute(id int64) error {
    return s.Repo.Delete(id)
}
