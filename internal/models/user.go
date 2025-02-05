package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);unique_index;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Phone     string `gorm:"type:varchar(20)"`
	Address   string `gorm:"type:text"`
	RoleID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Role Role `gorm:"foreignkey:RoleID"`
}

func (user *User) TableName() string {
	return "users"
}
