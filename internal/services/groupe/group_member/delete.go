package group_member

import repo "github.com/engrsakib/erp-system/internal/repository/group"

type DeleteMemberService struct {
    Repo *repo.MemberRepository
}

func NewDeleteMemberService(r *repo.MemberRepository) *DeleteMemberService {
    return &DeleteMemberService{Repo: r}
}

func (s *DeleteMemberService) Execute(id int64) error {
    return s.Repo.Delete(id)
}
