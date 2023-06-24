package services

import (
	"go-todolist/request"
)

type AuthService interface {
	Login(request request.LoginRequest) (string, error)
	Register(username string, email string, password string) (string, error)
}
