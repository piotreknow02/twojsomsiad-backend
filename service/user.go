package service

import (
	"twojsomsiad/model"

	"gorm.io/gorm"
)

func FindUserById(db *gorm.DB, id uint) (user model.User, err error) {
	err = db.Find(&user, id).Error
	return user, err
}

func UpdateUser(db *gorm.DB, id uint, data model.UserUpdateDTO) (newUser model.User, err error) {
	err = db.Find(&newUser, id).Error
	if err != nil {
		return model.User{}, err
	}
	if data.Name != "" {
		newUser.Name = data.Name
	}
	if data.Password != "" {
		newUser.Password, err = HashPassword(data.Password)
		if err != nil {
			return model.User{}, err
		}
	}
	if data.Surname != "" {
		newUser.Surname = data.Surname
	}
	if data.Username != "" {
		newUser.Username = data.Username
	}
	err = db.Save(&newUser).Error
	return newUser, err
}
