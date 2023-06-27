package repository

import (
	"go-todolist/model"
	"gorm.io/gorm"
)

type todoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepositoryImpl{db}
}

func (t *todoRepositoryImpl) GetByUserId(userId uint) (todos []model.Todo) {
	var newTodos []model.Todo
	err := t.db.Where("user_id = ?", userId).Find(&todos).Error
	if err != nil {
		return newTodos
	}

	return todos
}

func (t *todoRepositoryImpl) Create(userId uint, todo *model.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (t *todoRepositoryImpl) Update(userId uint, todo *model.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (t *todoRepositoryImpl) Delete(todoId int) error {
	//TODO implement me
	panic("implement me")
}
