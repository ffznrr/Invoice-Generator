package routes

import (
	"invoice_gen_be/internal/handler"
	"invoice_gen_be/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupAPIRoutes(api fiber.Router) {

	api.Post("/login", handler.Login)
	api.Get("/items", handler.GetItemsByCode)

	protected := api.Group("/", middleware.JWTProtected())

	protected.Post("/invoices", handler.SubmitInvoice)

}