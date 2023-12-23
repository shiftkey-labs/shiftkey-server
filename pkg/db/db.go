package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shiftkey-labs/shiftkey-server/pkg/data"
	"github.com/shiftkey-labs/shiftkey-server/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var err error
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")
	pgHost := os.Getenv("DB_HOST")
	pgPort := os.Getenv("DB_PORT")
	fmt.Println(pgUser)

	if pgHost == "" {
		pgHost = "localhost"
	}
	if pgPort == "" {
		pgPort = "5432"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", pgHost, pgUser, pgPass, pgDB, pgPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected")
}

func Migrate() {
	// For Role Type
	DB.Exec(`DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role') THEN
        CREATE TYPE Role AS ENUM ('student', 'volunteer', 'admin', 'god');
    END IF;
	END$$;`)

	// For EventStatus Type
	DB.Exec(`DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'eventstatus') THEN
        CREATE TYPE EventStatus AS ENUM ('draft', 'published', 'finished');
    END IF;
	END$$;`)

	err := DB.AutoMigrate(&model.User{}, &model.Host{}, &model.Event{})

	for _, user := range data.GetLocalUsers() {
		err := DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&user).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, event := range data.GetLocalEvents() {
		err := DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&event).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, host := range data.GetLocalHosts() {
		err := DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&host).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	DB.AutoMigrate(&model.EventHost{}, &model.Attendance{})

	for _, attendance := range data.GetLocalAttendances() {
		err := DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&attendance).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	for _, eventHost := range data.GetLocalEventHosts() {
		err := DB.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&eventHost).Error
		if err != nil {
			fmt.Println(err)
		}
	}
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}
