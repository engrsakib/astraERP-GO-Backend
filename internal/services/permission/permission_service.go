package permission

import (
	dto "github.com/engrsakib/erp-system/internal/dto/user/permission"
	permissionRepo "github.com/engrsakib/erp-system/internal/repository/permission"
)

type PermissionService struct {
	Repo *permissionRepo.PermissionRepository
}

func NewPermissionService(repo *permissionRepo.PermissionRepository) *PermissionService {
	return &PermissionService{Repo: repo}
}

func (s *PermissionService) AssignPermissions(req dto.AssignPermissionRequest) error {
	
	return s.Repo.SyncPermissions(req.TargetUserID, req.Permissions)
}