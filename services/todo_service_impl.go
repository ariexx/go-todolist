package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-todolist/model"
	"go-todolist/repository"
	"go-todolist/request"
)

type todoServiceImpl struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoServiceImpl{todoRepository}
}

func (t *todoServiceImpl) GetByUserId(userId uint) ([]model.Todo, error) {
	fmt.Println(userId)
	todos := t.todoRepository.GetByUserId(userId)
	return todos, nil
}

func (t *todoServiceImpl) Create(ctx *fiber.Ctx, userId uint, todo *request.CreateTodoRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoServiceImpl) Update(ctx *fiber.Ctx, userId uint, todoId int, todo *request.UpdateTodoRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoServiceImpl) Delete(ctx *fiber.Ctx, userId uint, todoId int) (string, error) {
	//TODO implement me
	panic("implement me")
}
