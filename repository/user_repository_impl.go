package repository

import (
	"errors"
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
		return user, errors.New("user not found")
	}
	return user, nil
}

func (repository *userRepositoryImpl) Create(user model.User) (model.User, error) {

	//check data given is exist or not
	if err := repository.db.Where("email = ?", user.Email).Where("username = ?", user.Username).
		First(&user).Error; err == nil {
		//handle custom error, if email or username already exist
		//to use custom error in golang, use errors.New("your message")
		//or you can use errors.Errorf("your message %s", variable)
		//im recommend to use custom error
		return user, errors.New("email or username already exist")
	}

	err := repository.db.Create(&user).Error
	if err != nil {
		return user, errors.New("failed to create user")
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := repository.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, errors.New("email not found")
	}
	return user, nil
}
