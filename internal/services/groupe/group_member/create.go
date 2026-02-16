package group_member

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    "github.com/engrsakib/erp-system/internal/models"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type CreateMemberService struct {
    Repo *repo.MemberRepository
}

func NewCreateMemberService(r *repo.MemberRepository) *CreateMemberService {
    return &CreateMemberService{Repo: r}
}

func (s *CreateMemberService) Execute(req group.MemberRequest, addedBy int64) (*models.Member, error) {
    member := &models.Member{
        UserID:  req.UserID,
        GroupID: req.GroupID,
        Status:  req.Status,
        AddedBy: addedBy,
    }

    err := s.Repo.Create(member)
    return member, err
}
