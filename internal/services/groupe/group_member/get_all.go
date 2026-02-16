package group_member

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    repo "github.com/engrsakib/erp-system/internal/repository/group"
)

type GetAllMemberService struct {
    Repo *repo.MemberRepository
}

func NewGetAllMemberService(r *repo.MemberRepository) *GetAllMemberService {
    return &GetAllMemberService{Repo: r}
}

func (s *GetAllMemberService) Execute(q group.MemberSearchQuery) (interface{}, interface{}, error) {
    if q.Page == 0 {
        q.Page = 1
    }
    if q.Limit == 0 {
        q.Limit = 10
    }

    data, total, err := s.Repo.GetAll(q.Page, q.Limit, q.Search)
    if err != nil {
        return nil, nil, err
    }

    meta := map[string]interface{}{
        "page":  q.Page,
        "limit": q.Limit,
        "total": total,
    }

    return data, meta, nil
}
