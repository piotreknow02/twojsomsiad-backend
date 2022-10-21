package service

import (
	"fmt"
	"twojsomsiad/model"

	"gorm.io/gorm"
)

func FindAdvertById(db *gorm.DB, id string) (advert model.Advert, err error) {
	err = db.First(&advert, id).Error
	fmt.Print("this is advert: 				", advert)
	return advert, err
}

func CreateAdvert(db *gorm.DB, data *model.CreateAdvertDTO) error {
	newAdvert := model.Advert{
		Title:       data.Title,
		Description: data.Description,
		City:        data.City,
		Date:        data.Date,
		UserID:      data.UserID,
	}
	return db.Create(&newAdvert).Error
}
