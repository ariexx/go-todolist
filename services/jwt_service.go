package services

type JwtService interface {
	GenerateToken(userId uint) (string, error)
}
