package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-todolist/config"
)

type jwtServiceImpl struct {
}

func (j *jwtServiceImpl) ValidateToken(token string) (*jwt.Token, error) {
	var secretKey = []byte(config.JwtSecret)
	newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	return newToken, nil
}

func NewJwtService() JwtService {
	return &jwtServiceImpl{}
}

func (j *jwtServiceImpl) GenerateToken(userId uint) (string, error) {
	var secretKey = []byte(config.JwtSecret)
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedString, nil
}
