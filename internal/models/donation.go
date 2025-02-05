package models

import (
	"time"
)

type Donation struct {
	ID             uint      `gorm:"primary_key"`
	UserID         uint      `gorm:"not null"`
	OrganizationID uint      `gorm:"not null"`
	Amount         float64   `gorm:"type:decimal(10,2);not null"`
	DonationDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status         string    `gorm:"type:varchar(20);default:'pending'"`
	PaymentMethod  string    `gorm:"type:varchar(50)"`
	TransactionID  string    `gorm:"type:varchar(255);unique_index"`

	User         User         `gorm:"foreignkey:UserID"`
	Organization Organization `gorm:"foreignkey:OrganizationID"`
}

func (donation *Donation) TableName() string {
	return "donations"
}
