package services

import (
	"go-todolist/request"
)

type AuthService interface {
	Login(request request.LoginRequest) (string, error)
	Register(request request.RegisterRequest) (string, error)
}
