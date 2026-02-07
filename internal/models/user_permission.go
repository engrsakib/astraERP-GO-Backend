package models

import "time"

type UserPermission struct {
	ID             uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         int64  `gorm:"not null;index" json:"user_id"`            
	PermissionSlug string `gorm:"type:varchar(100);not null;index" json:"permission_slug"` 

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CreatedAt time.Time `json:"created_at"`
}