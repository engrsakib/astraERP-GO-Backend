package user

import (
	
	dto "github.com/engrsakib/erp-system/internal/dto/user"
)

// UpdateUser Service
func (s *UserService) UpdateUser(id string, req dto.UpdateUserRequest) (*dto.UserResponse, error) {
    
    user, err := s.Repo.GetUserByID(id)
    if err != nil {
        return nil, err
    }

    
    user.Name = req.Name
    user.Email = req.Email
    
    
   
    if req.UserType != 0 {
        user.UserType = req.UserType
    }
    
    user.Photo = req.Photo

 
    if err := s.Repo.UpdateUser(user); err != nil {
        return nil, err
    }

    
    response := &dto.UserResponse{
        ID:        user.ID,
        Name:      user.Name,
        Mobile:    user.Mobile, 
        Email:     user.Email,
        UserType:  user.UserType,
        Photo:     user.Photo,
        CreatedAt: user.CreatedAt,
    }

    return response, nil
}
