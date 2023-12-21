package controller

import (
	"github.com/shiftkey-labs/shiftkey-api/pkg/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	res := DB.Find(&users)
	return users, res.Error
}
