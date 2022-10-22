package database

import (
	"fmt"

	"twojsomsiad/config"
	"twojsomsiad/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	DbErr error
)

func GetDB() *gorm.DB {
	return DB
}

func Setup() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw",
		config.Conf.DBHost,
		config.Conf.DBUser,
		config.Conf.DBPassword,
		"twojsomsiad",
		"5432",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		DbErr = err
		return err
	}

	db.AutoMigrate(
		&model.User{},
		&model.Advert{},
		&model.Application{},
	)

	DB = db

	return nil
}
