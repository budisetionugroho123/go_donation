package models

import (
	"time"
)

type Organization struct {
	ID          uint   `gorm:"primary_key"`
	UserID      uint   `gorm:"not null"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	LogoURL     string `gorm:"type:varchar(255)"`
	ContactInfo string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User User `gorm:"foreignkey:UserID"`
}

func (organization *Organization) TableName() string {
	return "organizations"
}
