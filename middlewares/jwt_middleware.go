package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
			//ini mengambil Bearer dari token (Bearer <token>)
			// kenapa 7? karena Bearer ada 6 karakter dan ada 1 spasi
			jwtToken = jwtToken[7:]

			//validate token with jwt service
			jwtService := services.NewJwtService()
			token, err := jwtService.ValidateToken(jwtToken)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFail("Unauthorized", "Unauthorized", nil))
			}

			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFail("Unauthorized", "Unauthorized", nil))
			}

			userId := uint(claim["user_id"].(float64))
			//set user_id to context
			c.Locals("user_id", userId)
			return c.Next()
		}
	}
}
