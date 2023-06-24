package repository

import (
	"go-todolist/model"
)

type UserRepository interface {
	GetUserById(id int) (model.User, error)
	Create(user model.User) (model.User, error)
}
