package models

type GroupType struct {
    ID   int64  `gorm:"primaryKey;autoIncrement" json:"id"`
    Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func (GroupType) TableName() string {
    return "group_types"
}
