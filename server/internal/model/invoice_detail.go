package model

type InvoiceDetail struct {
	ID        uint `gorm:"primaryKey"`
	InvoiceID uint
	ItemID    uint
	Quantity  int
	Price     int
	Subtotal  int

	// optional relation (biar bisa preload)
	Item Item `gorm:"foreignKey:ItemID"`
}