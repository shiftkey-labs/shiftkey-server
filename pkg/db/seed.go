package db

import (
	"github.com/shiftkey-labs/shiftkey-api/pkg/data"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(db *gorm.DB) error {
	for _, user := range data.GetLocalUsers() {
		err := db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&user).Error
		if err != nil {
			return err
		}
	}
	return nil
}
