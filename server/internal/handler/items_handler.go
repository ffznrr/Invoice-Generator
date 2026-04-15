package handler

import (
	"invoice_gen_be/internal/database"
	"invoice_gen_be/internal/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetItemsByCode(c *fiber.Ctx) error {
	code := c.Query("code")

	if code == "" {
		return c.JSON([]model.Item{})
	}

	var items []model.Item

	err := database.DB.
		Where("LOWER(code) LIKE ?", "%"+strings.ToLower(code)+"%").
		Limit(20).
		Find(&items).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to fetch items",
			"error":   err.Error(),
		})
	}

	return c.JSON(items)
}