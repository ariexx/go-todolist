package repository

import (
	"go-todolist/model"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (repository *userRepositoryImpl) GetUserById(id int) (model.User, error) {
	var user model.User
	err := repository.db.First(&user, id).Error
	//check if user not found
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) Create(user model.User) (model.User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
