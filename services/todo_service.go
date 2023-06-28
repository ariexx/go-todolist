package services

import (
	"github.com/gofiber/fiber/v2"
	"go-todolist/model"
	"go-todolist/request"
)

type TodoService interface {
	GetByUserId(userId uint) ([]model.Todo, error)
	Create(userId uint, todo *request.CreateTodoRequest) (string, error)
	Update(ctx *fiber.Ctx, userId uint, todoId int, todo *request.UpdateTodoRequest) (string, error)
	Delete(ctx *fiber.Ctx, userId uint, todoId int) (string, error)
}
