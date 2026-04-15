package handler

import (
	"time"

	"invoice_gen_be/internal/config"
	"invoice_gen_be/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if(req.Username == ""){
		return c.Status(401).JSON(fiber.Map{"error": "Silahkan Isi Username terlebih dahulu"})
	}

	if(req.Password == ""){
		return c.Status(401).JSON(fiber.Map{"error": "Silahkan Isi Password terlebih dahulu"})
	}
	
	user, ok := service.Authenticate(req.Username, req.Password)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Username atau login salah"})
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"username": user.Username,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString(config.JWT_SECRET)

	return c.JSON(fiber.Map{
		"token": t,
	})
}