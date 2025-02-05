package models

import (
	"time"
)

type Transaction struct {
	ID                 uint      `gorm:"primary_key"`
	DonationID         uint      `gorm:"not null"`
	Amount             float64   `gorm:"type:decimal(10,2);not null"`
	PaymentStatus      string    `gorm:"type:varchar(20);default:'pending'"`
	PaymentDate        time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	PaymentMethod      string    `gorm:"type:varchar(50)"`
	TransactionDetails string    `gorm:"type:text"`

	Donation Donation `gorm:"foreignkey:DonationID"`
}

func (transaction *Transaction) TableName() string {
	return "transactions"
}
