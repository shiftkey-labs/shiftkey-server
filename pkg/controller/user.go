package controller

import (
	"github.com/shiftkey-labs/shiftkey-api/pkg/db"
	"github.com/shiftkey-labs/shiftkey-api/pkg/model"
)

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	res := db.DB.Find(&users)
	return users, res.Error
}
