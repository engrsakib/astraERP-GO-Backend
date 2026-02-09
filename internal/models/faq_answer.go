package models

import "time"

type FaqAnswer struct {
	ID        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	FaqID     int64  `gorm:"not null;index" json:"faq_id"` 
	Question  string `gorm:"type:text;not null" json:"question"`
	Answer    string `gorm:"type:text;not null" json:"answer"`

	AddedBy   *int64 `json:"added_by"` 

	// Relations
	Faq       Faq    `gorm:"foreignKey:FaqID" json:"-"`

	
	User      *User  `gorm:"foreignKey:AddedBy;constraint:OnDelete:SET NULL" json:"-"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}