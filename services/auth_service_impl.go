package services

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"go-todolist/helper"
	"go-todolist/model"
	"go-todolist/repository"
	"go-todolist/request"
)

type authServiceImpl struct {
	repository repository.UserRepository
	jwtService JwtService
}

func NewAuthService(repository repository.UserRepository, jwtService JwtService) AuthService {
	return &authServiceImpl{repository, jwtService}
}

func (service *authServiceImpl) Login(request request.LoginRequest) (string, error) {
	err := validator.New().Struct(&request)
	if err != nil {
		return "", err
	}
	user, err := service.repository.FindByEmail(request.Email)
	if err != nil {
		return "", err
	}
	if user.Password != request.Password {
		return "", errors.New("email or password is wrong")
	}

	token, err := service.jwtService.GenerateToken(uint(user.Id))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (service *authServiceImpl) Register(request request.RegisterRequest) (string, error) {
	err := validator.New().Struct(&request)
	if err != nil {
		return "", err
	}

	user := model.User{}
	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password
	user.CreatedAt = helper.Now()
	user.UpdatedAt = helper.Now()

	create, err := service.repository.Create(user)
	if err != nil {
		return "", err
	}

	//generate jwt token
	token, err := service.jwtService.GenerateToken(uint(create.Id))
	if err != nil {
		return "", err
	}

	return token, nil
}
