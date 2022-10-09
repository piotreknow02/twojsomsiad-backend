package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" validate:"required" gorm:"unique"`
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"bcrypt"`
	Points   uint   `json:"points"`
}
