package service

import (
	"twojsomsiad/config"
	"twojsomsiad/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FindUserByCredentials(db *gorm.DB, credentials *model.AuthLoginDTO) (authorized bool, user model.User) {
	db.Find(&user, "email = ?", credentials.Email)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err == nil {
		return true, user
	}
	return false, model.User{}
}

func CreateUser(db *gorm.DB, data *model.AuthRegisterDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := model.User{
		Username: data.Username,
		Name:     data.Name,
		Surname:  data.Surname,
		Email:    data.Email,
		Password: string(hashedPassword),
		Points:   config.Conf.DefaultPoints,
	}
	return db.Create(&newUser).Error
}
