package entities

import "time"

type Category struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement;"`
	CategoryName string    `gorm:"type:varchar(255)" json:"category_name"`
	Description  string    `gorm:"type:varchar(255)" json:"description"`
	Status       string    `gorm:"type:varchar(255)" json:"status"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;"`
}
