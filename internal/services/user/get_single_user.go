package user


import (
 
	dto "github.com/engrsakib/erp-system/internal/dto/user"
)	


// GetUser: সার্ভিস মেথড
func (s *UserService) GetUser(id string) (*dto.UserResponse, error) {
	// ১. রিপোজিটরি থেকে ডাটা আনা
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// ২. পার্মিশন অবজেক্ট থেকে শুধু নামগুলো বের করে লিস্ট বানানো
	var permissionList []string
	
	// আপনার models.UserPermission স্ট্রাক্টে যেই ফিল্ডে পার্মিশন নাম আছে (ধরি 'Permission')
	// সেটা লুপ চালিয়ে বের করতে হবে
	for _, p := range user.Permissions {
		// ⚠️ সতর্কতা: আপনার মডেলে এই ফিল্ডের নাম যা আছে (PermissionName বা Name), তাই দিন।
		// আমি ধরে নিচ্ছি এটা 'PermissionName' হবে।
		permissionList = append(permissionList, p.PermissionSlug) 
	}

	// ৩. রেসপন্স তৈরি করা
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