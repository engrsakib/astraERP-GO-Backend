package user


import (
 
	dto "github.com/engrsakib/erp-system/internal/dto/user"
)	


func (s *UserService) GetUser(id string) (*dto.UserResponse, error) {
	
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	
	var permissionList []string
	

	for _, p := range user.Permissions {
		
		permissionList = append(permissionList, p.PermissionSlug) 
	}

	response := &dto.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Mobile:      user.Mobile,
		Email:       user.Email,
		UserType:    user.UserType, 
		Photo:       user.Photo,
		Permissions: permissionList,
		CreatedAt:   user.CreatedAt,
	}

	return response, nil
}