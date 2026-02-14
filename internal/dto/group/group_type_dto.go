// package group

// type CreateGroupTypeRequest struct {
// 	Name string `json:"name" binding:"required,min=2,max=100"`
// }

// type UpdateGroupTypeRequest struct {
// 	Name string `json:"name" binding:"required,min=2,max=100"`
// }


package group

type GroupTypeRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
}
