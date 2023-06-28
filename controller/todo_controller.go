package controller

import "github.com/gofiber/fiber/v2"

type TodoController interface {
	GetByUserId(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(userId uint, todoId int, todo any) error
	Delete(userId uint, todoId int) error
}
