package models

import "time"

type User struct {
    ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Name      string    `gorm:"type:varchar(100);not null" json:"name"`
    Mobile    string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"mobile"`
    Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
    Password  string    `gorm:"type:varchar(255);not null" json:"-"`
    Photo     string    `gorm:"type:varchar(255)" json:"photo"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
