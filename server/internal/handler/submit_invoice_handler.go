package handler

import (
	"invoice_gen_be/internal/database"
	"invoice_gen_be/internal/dto"
	"invoice_gen_be/internal/model"
	"log"
	"time"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

func generateInvoiceNumber() string {
	return fmt.Sprintf("INV-%d", time.Now().UnixNano())
}
func SubmitInvoice(c *fiber.Ctx) error {
	var req dto.SubmitInvoiceRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"error":   err.Error(),
		})
	}


	username, ok := c.Locals("username").(string)
	log.Println(username)
	log.Println("username")

	if !ok || username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized user",
		})
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to start transaction",
		})
	}

	// auto rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var (
		totalAmount int64
		details     []model.InvoiceDetail
	)

	// create invoice header
	invoice := model.Invoice{
		InvoiceNumber:   generateInvoiceNumber(),
		SenderName:      req.SenderName,
		SenderAddress:   req.SenderAddress,
		ReceiverName:    req.ReceiverName,
		ReceiverAddress: req.ReceiverAddress,
		CreatedBy:       username,
		CreatedAt:       time.Now(),
	}

	if err := tx.Create(&invoice).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create invoice",
			"error":   err.Error(),
		})
	}

	// ZERO TRUST VALIDATION
	for _, itemReq := range req.Items {
		var item model.Item

		if err := tx.First(&item, itemReq.ItemID).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "item not found",
				"item_id": itemReq.ItemID,
			})
		}

		if itemReq.Quantity <= 0 {
			tx.Rollback()
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid quantity",
				"item_id": itemReq.ItemID,
			})
		}

		subtotal := item.Price * itemReq.Quantity
		totalAmount += int64(subtotal)

		details = append(details, model.InvoiceDetail{
			InvoiceID: invoice.ID,
			ItemID:    item.ID,
			Quantity:  itemReq.Quantity,
			Price:     item.Price,
			Subtotal:  subtotal,
		})
	}

	// insert details batch
	if len(details) > 0 {
		if err := tx.Create(&details).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to create invoice details",
				"error":   err.Error(),
			})
		}
	}

	// update total
	if err := tx.Model(&model.Invoice{}).
		Where("id = ?", invoice.ID).
		Update("total_amount", totalAmount).Error; err != nil {

		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update total amount",
			"error":   err.Error(),
		})
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to commit transaction",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "invoice created successfully",
		"data":    invoice,
	})
}