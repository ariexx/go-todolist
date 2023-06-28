package services

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
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

func (t *todoServiceImpl) Create(userId uint, todo *request.CreateTodoRequest) (string, error) {
	err := validator.New().Struct(todo)
	if err != nil {
		return "", errors.New("validation failed")
	}

	todoData := model.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		UserId:      userId,
	}

	result, err := t.todoRepository.Create(userId, &todoData)
	if err != nil {
		return "", errors.New("failed to create todo")
	}

	if !result {
		return "", errors.New("failed to create todo")
	}

	return "success", nil
}

func (t *todoServiceImpl) Update(ctx *fiber.Ctx, userId uint, todoId int, todo *request.UpdateTodoRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (t *todoServiceImpl) Delete(ctx *fiber.Ctx, userId uint, todoId int) (string, error) {
	//TODO implement me
	panic("implement me")
}
