package models

import (
	"time"
)

type Group struct {
	ID           int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string     `gorm:"type:varchar(255);not null" json:"name"`
	
	
	GroupTypeID  int64      `json:"group_type_id"`
	GroupType    GroupType  `gorm:"foreignKey:GroupTypeID" json:"group_type"`

	
	CheckInTime  *time.Time `gorm:"type:time" json:"check_in_time"`
	CheckOutTime *time.Time `gorm:"type:time" json:"check_out_time"`

	AddedBy      int64      `json:"added_by"`
	Creator      User       `gorm:"foreignKey:AddedBy" json:"creator"`

	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (Group) TableName() string {
	return "groups"
}