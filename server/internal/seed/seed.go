package seed

import (
	"invoice_gen_be/internal/database"
	"invoice_gen_be/internal/model"
	"log"
)

func SeedItems() {
	var count int64
	database.DB.Model(&model.Item{}).Count(&count)

	if count > 0 {
		return
	}

	items := []model.Item{
		{Code: "BRG-001", Name: "Pensil", Price: 2000},
		{Code: "BRG-002", Name: "Pulpen", Price: 3000},
		{Code: "BRG-003", Name: "Buku Tulis", Price: 5000},
	}

	if err := database.DB.Create(&items).Error; err != nil {
		log.Fatal("Seed failed:", err)
	}

	log.Println("Seeder items success")
}