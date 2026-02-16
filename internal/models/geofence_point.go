package models

import (
	"time"
	"gorm.io/gorm"
)

type GeofencePoint struct {
    ID         int64          `gorm:"primaryKey" json:"id"`
    GeofenceID int64          `gorm:"not null;index" json:"geofence_id"`
    Latitude   float64        `gorm:"type:decimal(10,8);not null" json:"latitude"`
    Longitude  float64        `gorm:"type:decimal(11,8);not null" json:"longitude"`
    
    
    Geom       string         `gorm:"type:geometry(Point,4326)" json:"-"` 
    
    SeqOrder   int            `gorm:"not null;default:0" json:"seq_order"`
    CreatedAt  time.Time      `json:"created_at"`
    UpdatedAt  time.Time      `json:"updated_at"`
    DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}