package database

import (
	"invoice_gen_be/internal/model"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(
		&model.Item{},
		&model.Invoice{},
		&model.InvoiceDetail{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("✅ Migration success")
}