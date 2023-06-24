package middlewares

import "github.com/gofiber/fiber/v2"

func GeneralMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Set("Accept", "application/json")
		c.Set("Content-Type", "application/json")
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, Token")
		c.Set("Access-Control-Allow-Credentials", "true")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(200)
		}
		return c.Next()
	}
}
