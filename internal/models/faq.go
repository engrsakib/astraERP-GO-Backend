package models

import "time"

type Faq struct {
    ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
    Photo     string    `gorm:"type:varchar(255)" json:"photo"`
    Headline  string    `gorm:"type:varchar(255);not null" json:"headline"`

    
    AddedBy   *int64    `json:"added_by"`

    // FAQ → FAQ Answers (One-to-Many)
    FaqAnswers []FaqAnswer `gorm:"foreignKey:FaqID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"faq_answers"`

    // FAQ → User (Belongs To)
    User      *User     `gorm:"foreignKey:AddedBy;constraint:OnDelete:SET NULL" json:"user"`

    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
