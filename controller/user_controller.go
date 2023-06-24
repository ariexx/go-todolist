package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	GetUserById(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
}
