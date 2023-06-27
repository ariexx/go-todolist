package repository

import "go-todolist/model"

type TodoRepository interface {
	GetByUserId(userId uint) (todos []model.Todo)
	Create(userId uint, todo *model.Todo) error
	Update(userId uint, todo *model.Todo) error
	Delete(todoId int) error
}
