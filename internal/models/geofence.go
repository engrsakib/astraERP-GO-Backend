package models

import (
	"time"
	"gorm.io/gorm"
)


type Geofence struct {
    ID          int64          `gorm:"primaryKey" json:"id"`
    Name        string         `gorm:"type:varchar(255);not null;index" json:"name"`
    Description string         `gorm:"type:text" json:"description"`
    Points      []GeofencePoint `gorm:"foreignKey:GeofenceID" json:"points"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // সফট ডিলিট
}