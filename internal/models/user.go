package models

import "time"

const (
    UserTypeSuperAdmin = 0 // super Admin: most powerful user
    UserTypeManager    = 1 // Manager: 
    UserTypeAdmin      = 2 // Admin (Subscriber): subcriber level admin
    UserTypeUser       = 3 // Regular User: default user type
)

type User struct {
    ID       int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Name     string    `gorm:"type:varchar(100);not null" json:"name"`
    Mobile   string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"mobile"`
    Email    string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
    Password string    `gorm:"type:varchar(255);not null" json:"-"`
    Photo    string    `gorm:"type:varchar(255)" json:"photo"`

    
    UserType int8 `gorm:"type:smallint;default:3;not null" json:"user_type"` 
    
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}