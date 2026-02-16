package group_member

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type UpdateMemberService struct {
    Repo *repo.MemberRepository
}

func NewUpdateMemberService(r *repo.MemberRepository) *UpdateMemberService {
    return &UpdateMemberService{Repo: r}
}

func (s *UpdateMemberService) Execute(id int64, req group.MemberRequest) error {
    member, err := s.Repo.GetByID(id)
    if err != nil {
        return err
    }

    member.UserID = req.UserID
    member.GroupID = req.GroupID
    member.Status = req.Status

    return s.Repo.Update(member)
}
