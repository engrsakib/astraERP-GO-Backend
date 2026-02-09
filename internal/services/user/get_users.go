package user

import (
	"math"

	dto "github.com/engrsakib/erp-system/internal/dto/user"
	
)

// GetUsers সার্ভিস মেথড (UserService স্ট্রাক্টের অংশ)
// আপনার user_service.go ফাইলে UserService স্ট্রাক্টটি থাকতে হবে
func (s *UserService) GetUsers(query dto.PaginationQuery) ([]dto.UserResponse, *dto.PaginationMeta, error) {
	
	// ১. রিপোজিটরি কল
	users, total, err := s.Repo.GetUsers(query.Page, query.Limit, query.Search)
	if err != nil {
		return nil, nil, err
	}

	// ২. মডেল থেকে DTO তে কনভার্ট
	var userResponses []dto.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Mobile:    u.Mobile,
			Email:     u.Email,
			UserType:  u.UserType, // টাইপ কাস্টিং লাগলে করবেন (যেমন int(u.UserType))
			Photo:     u.Photo,
			CreatedAt: u.CreatedAt,
		})
	}

	// ৩. মেটা ডাটা তৈরি
	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))
	
	meta := &dto.PaginationMeta{
		CurrentPage: query.Page,
		TotalPages:  totalPages,
		TotalItems:  total,
		Limit:       query.Limit,
	}

	return userResponses, meta, nil
}