package model

type Item struct {
	ID    uint   `gorm:"primaryKey"`
	Code  string `gorm:"unique;not null"`
	Name  string
	Price int
}