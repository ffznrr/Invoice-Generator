package routes

import (
	"invoice_gen_be/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupAPIRoutes(api fiber.Router) {

	api.Post("/login", handler.Login)
	api.Get("/items", )

}