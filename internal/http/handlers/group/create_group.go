package group

type GroupRequest struct {
    Name         string `json:"name" binding:"required"`
    GroupTypeID  int64  `json:"group_type_id" binding:"required"`
    CheckInTime  string `json:"check_in_time"`
    CheckOutTime string `json:"check_out_time"`
}
