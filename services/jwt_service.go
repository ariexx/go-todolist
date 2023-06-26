package services

import "github.com/golang-jwt/jwt/v5"

type JwtService interface {
	GenerateToken(userId uint) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
