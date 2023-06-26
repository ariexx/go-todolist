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
		return user, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) Create(user model.User) (model.User, error) {

	//check data given is exist or not
	queryCheck := repository.db.Where("email = ?", user.Email).Where("username = ?", user.Username).First(&user)
	if queryCheck.RowsAffected != 0 {
		return user, queryCheck.AddError(errors.New("data already exist"))
	}

	err := repository.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repository *userRepositoryImpl) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := repository.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
