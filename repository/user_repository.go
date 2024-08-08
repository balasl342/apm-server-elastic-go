package repository

import (
	"github.com/balasl342/apm-server-elastic-go/database"
	"github.com/balasl342/apm-server-elastic-go/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id int) (*models.User, error) {
	var user models.User
	res := database.DB.First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	return database.DB.Create(&user).Error
}
