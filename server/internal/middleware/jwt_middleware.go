package middleware

import (
	"log"
	"strings"

	"invoice_gen_be/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "UNAUTHORIZED"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "UNAUTHORIZED"})
		}

		userID := claims["user_id"]
		username := claims["username"]

	log.Println(" middleware username")
	log.Println(username)
	log.Println(" middleware username")


		c.Locals("user_id", userID)
		c.Locals("username", username)

		return c.Next()
	}
}