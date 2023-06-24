package services

import (
	"github.com/go-playground/validator/v10"
	"go-todolist/model"
	"go-todolist/repository"
	"go-todolist/request"
	"time"
)

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userServiceImpl{repository}
}

func (service *userServiceImpl) GetUserById(id int) (model.User, error) {
	return service.repository.GetUserById(id)
}

func (service *userServiceImpl) CreateUser(request request.CreateUserRequest) (model.User, error) {
	//implement validation
	err := validator.New().Struct(&request)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{}
	user.Username = request.Username
	user.Password = request.Password
	user.Email = request.Email
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	newUser, err := service.repository.Create(user)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
