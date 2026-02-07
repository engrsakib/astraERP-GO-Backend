package permission

type AssignPermissionRequest struct {
	TargetUserID int64    `json:"user_id" binding:"required"` 
	Permissions  []string `json:"permissions" binding:"required"` 
}