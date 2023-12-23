package db

import (
	"fmt"

	"github.com/shiftkey-labs/shiftkey-server/pkg/data"
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

	for _, event := range data.GetLocalEvents() {
		err := db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&event).Error
		if err != nil {
			return err
		}
	}

	for _, host := range data.GetLocalHosts() {
		err := db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&host).Error
		if err != nil {
			return err
		}
	}

	for _, attendance := range data.GetLocalAttendances() {
		err := db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&attendance).Error
		if err != nil {
			return err
		}
	}

	for _, eventHost := range data.GetLocalEventHosts() {
		err := db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&eventHost).Error
		if err != nil {
			return err
		}
	}

	fmt.Println("Database seeded")

	return nil
}
