package group_member

import repo "github.com/engrsakib/erp-system/internal/repository/group"

type GetMemberByIDService struct {
    Repo *repo.MemberRepository
}

func NewGetMemberByIDService(r *repo.MemberRepository) *GetMemberByIDService {
    return &GetMemberByIDService{Repo: r}
}

func (s *GetMemberByIDService) Execute(id int64) (interface{}, error) {
    return s.Repo.GetByID(id)
}
