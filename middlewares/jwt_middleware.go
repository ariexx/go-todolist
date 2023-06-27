package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go-todolist/helper"
	"go-todolist/services"
)

func JwtMiddlewares() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jwtToken := c.Get("Authorization")
		switch jwtToken {
		case "":
			return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFail("Unauthorized", "Unauthorized", nil))
		case "Bearer ": // empty token
			return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFail("Unauthorized", "Unauthorized", nil))
		default:
			//get token from header
			jwtToken = jwtToken[7:]

			//validate token with jwt service
			jwtService := services.NewJwtService()
			_, err := jwtService.ValidateToken(jwtToken)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFail("Unauthorized", "Unauthorized", nil))
			}
			return c.Next()

		}
	}
}
