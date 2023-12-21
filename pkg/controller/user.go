package controller

import (
	"github.com/shiftkey-labs/shiftkey-server/pkg/db"
	"github.com/shiftkey-labs/shiftkey-server/pkg/model"
)

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	res := db.DB.Find(&users)
	return users, res.Error
}
