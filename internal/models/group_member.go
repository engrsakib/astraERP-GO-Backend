package models

import (
    "time"

    "gorm.io/gorm"
)

const (
    MemberStatusPending = 0
    MemberStatusActive  = 1
    MemberStatusBlocked = 2
)


type Member struct {
    ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`

    UserID    int64          `json:"user_id"`
    User      User           `gorm:"foreignKey:UserID" json:"user"`

    GroupID   int64          `json:"group_id"`
    Group     Group          `gorm:"foreignKey:GroupID" json:"group"`

    Status    int8           `json:"status"` // 0=pending, 1=active, 2=blocked etc.

    AddedBy   int64          `json:"added_by"`
    Creator   User           `gorm:"foreignKey:AddedBy" json:"creator"`

    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (Member) TableName() string {
    return "members"
}
