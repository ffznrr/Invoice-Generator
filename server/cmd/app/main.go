package main

import (
	"invoice_gen_be/internal/database"
	"invoice_gen_be/internal/routes"
	"invoice_gen_be/internal/seed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
	log.Println("No .env file found")
    }

	database.ConnectDB()

	app := fiber.New()

    seed.SeedItems()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server Is Running..")
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}