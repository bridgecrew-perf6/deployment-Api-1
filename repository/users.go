package repository

import (
	"deployment/models"

	"gorm.io/gorm"
)

type NewRepositoryUser interface {
	GetUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id int) error
	UpdateUser(user models.User) (models.User, error)
	Login(email, password string) (models.User, error)
}

type repositoryUser struct {
	db *gorm.DB
}

func InstenceRepoUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db}
}

func (r *repositoryUser) GetUsers() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repositoryUser) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repositoryUser) CreateUser(user models.User) (models.User, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (r *repositoryUser) DeleteUser(id int) error {
	var user models.User
	err := r.db.Delete(&user, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repositoryUser) UpdateUser(user models.User) (models.User, error) {

	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

// LOGIN
func (r *repositoryUser) Login(email, password string) (models.User, error) {
	foundUser := models.User{}
	if err := r.db.Where("email = ? AND password = ?", email, password).Find(&foundUser).Error; err != nil {
		return foundUser, err
	}
	return foundUser, nil
}

// // LOGIN
// func (r *repository) Login(email, password string) (models.User, error) {
// 	foundUser := models.User{}
// 	if err := r.db.Where("email = ? AND password = ?", email, password).Find(&foundUser).Error; err != nil {
// 		return foundUser, err
// 	}
// 	return foundUser, nil
// }
