package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB"),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	db.AutoMigrate(
		&Experience{},
		&User{},
		&Education{},
		&Certificate{},
		&Skill{},
		&Portofolio{},
		&ImageUrl{},
		&JobDetails{},
	)

	DB = db
}
