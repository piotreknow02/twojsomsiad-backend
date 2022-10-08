package database

import (
	"fmt"
	"twojsomsiad/config"

	"gorm.io/driver/mysql"
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
	connstring := config.Conf.DBUser + ":" + config.Conf.DBPassword + "@tcp(" + config.Conf.DBHost + ":3306)/locker?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(connstring)
	db, err := gorm.Open(
		mysql.Open(connstring),
		&gorm.Config{},
	)

	if err != nil {
		DbErr = err
		return err
	}

	// db.AutoMigrate(
	// 	// models
	// )

	DB = db

	return nil
}
