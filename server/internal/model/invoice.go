package model

import "time"

type Invoice struct {
	ID              uint             `gorm:"primaryKey"`
	InvoiceNumber   string           `gorm:"unique"`
	SenderName      string
	SenderAddress   string
	ReceiverName    string
	ReceiverAddress string
	TotalAmount     int
	CreatedBy       uint
	CreatedAt       time.Time

	Details []InvoiceDetail `gorm:"foreignKey:InvoiceID"`
}