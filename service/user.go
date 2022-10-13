package service

import (
	"twojsomsiad/model"

	"gorm.io/gorm"
)

func FindUserById(db *gorm.DB, id uint) (user model.User, err error) {
	err = db.Find(&user, id).Error
	return user, err
}

func UpdateUser(db *gorm.DB, id uint, data model.UserUpdateDTO) error {
	var newUser model.User
	err := db.Find(&newUser, id).Error
	if err != nil {
		return err
	}
	if data.Name != "" {
		newUser.Name = data.Name
	}
	if data.Password != "" {
		newUser.Password = data.Password
	}
	if data.Surname != "" {
		newUser.Surname = data.Surname
	}
	if data.Username != "" {
		newUser.Username = data.Username
	}
	return db.Save(&newUser).Error
}
