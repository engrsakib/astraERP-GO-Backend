package group

type GroupRequest struct {
    Name        string `json:"name"`
    GroupTypeID int64  `json:"group_type_id"`
}
