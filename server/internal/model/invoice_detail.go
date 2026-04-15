package model

type InvoiceDetail struct {
	ID        uint `gorm:"primaryKey"`
	InvoiceID uint
	ItemID    uint
	Quantity  int
	Price     int
	Subtotal  int

	Item Item `gorm:"foreignKey:ItemID"`
}