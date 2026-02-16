package models

import (
    "time"
    "gorm.io/gorm"
)

type GroupType struct {
    ID   int64  `gorm:"primaryKey;autoIncrement" json:"id"`
    Name string `gorm:"type:varchar(255);not null" json:"name"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (GroupType) TableName() string {
    return "group_types"
}
