package models

import "time"

type Role struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"type:varchar(255);unique" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (role *Role) TableName() string {
	return "roles"
}
