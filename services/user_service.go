package services

import (
	"go-todolist/model"
	"go-todolist/request"
)

type UserService interface {
	GetUserById(id int) (model.User, error)
	CreateUser(user request.CreateUserRequest) (model.User, error)
}
