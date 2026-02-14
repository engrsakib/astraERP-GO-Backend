package user

import (
	"math"

	dto "github.com/engrsakib/erp-system/internal/dto/user"
	
)


func (s *UserService) GetUsers(query dto.PaginationQuery) ([]dto.UserResponse, *dto.PaginationMeta, error) {
	
	
	users, total, err := s.Repo.GetUsers(query.Page, query.Limit, query.Search)
	if err != nil {
		return nil, nil, err
	}

	
	var userResponses []dto.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Mobile:    u.Mobile,
			Email:     u.Email,
			UserType:  u.UserType, 
			Photo:     u.Photo,
			CreatedAt: u.CreatedAt,
		})
	}


	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
	
	meta := &dto.PaginationMeta{
		CurrentPage: query.Page,
		TotalPages:  totalPages,
		TotalItems:  total,
		Limit:       query.Limit,
	}

	return userResponses, meta, nil
}