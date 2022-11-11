package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string   `json:"username" validate:"required" gorm:"unique"`
	Name     string   `json:"name" validate:"required"`
	Surname  string   `json:"surname" validate:"required"`
	Email    string   `json:"email" validate:"email"`
	Password string   `json:"password" validate:"bcrypt"`
	Points   uint     `json:"points"`
	Adverts  []Advert `json:"adverts"`
}

type UserUpdateDTO struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

type UserView struct {
	gorm.Model
	Username string   `json:"username" validate:"required" gorm:"unique"`
	Name     string   `json:"name" validate:"required"`
	Surname  string   `json:"surname" validate:"required"`
	Email    string   `json:"email" validate:"email"`
	Points   uint     `json:"points"`
	Adverts  []Advert `json:"adverts"`
}

func (user *User) ToView() UserView {
	return UserView{
		Model:    user.Model,
		Username: user.Username,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Points:   user.Points,
	}
}
