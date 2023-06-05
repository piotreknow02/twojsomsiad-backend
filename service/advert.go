package service

import (
	"errors"
	"twojsomsiad/model"

	"gorm.io/gorm"
)

func FindAdverts(db *gorm.DB, args *model.Args) (adverts []model.Advert, err error) {
	query := db.Select("adverts.*")
	query = query.Offset(args.Offset)
	query = query.Limit(args.Limit)
	if err := query.Preload("User").Find(&adverts).Error; err != nil {
		return []model.Advert{}, err
	}
	return adverts, nil
}

func FindAdvertById(db *gorm.DB, id string) (advert model.Advert, err error) {
	err = db.Preload("User").First(&advert, id).Error
	return advert, err
}

func CreateAdvert(db *gorm.DB, id uint, data *model.CreateAdvertDTO) (advert model.Advert, err error) {
	advert = model.Advert{
		Title:       data.Title,
		Description: data.Description,
		City:        data.City,
		Date:        data.Date,
		UserID:      id,
	}
	err = db.Create(&advert).Error
	return advert, err
}

func FindAdvertsForUser(db *gorm.DB, id uint) (adverts []model.Advert, err error) {
	err = db.Preload("User").Where("user_id = ?", id).Find(&adverts).Error
	return adverts, err
}

func CheckUserApplicationDuplicate(db *gorm.DB, userId uint, advertId uint) (exists bool, err error) {
	var adverts []model.Application
	queryErr := db.Where("advert_id = ?", advertId).Where("user_id = ?", userId).Find(&adverts).Error
	if queryErr == gorm.ErrRecordNotFound || len(adverts) < 1 {
		return false, nil
	}
	return true, queryErr
}

func ApplyForEvent(db *gorm.DB, userId uint, advertId uint) (application model.Application, err error) {
	exists, err := CheckUserApplicationDuplicate(db, userId, advertId)

	if err != nil {
		return model.Application{}, err
	}
	if exists {
		return model.Application{}, errors.New("duplicate advert application")
	}

	application = model.Application{
		UserID:   userId,
		AdvertID: advertId,
		Verified: false,
	}
	err = db.Create(&application).Error

	return application, err
}

func GetApplicationsForAdvert(db *gorm.DB, advertId uint) (applications []model.Application, err error) {
	err = db.Where("advert_id = ?", advertId).Find(&applications).Error
	return applications, err
}

func ConfirmApplication(db *gorm.DB, applicationId uint, userId uint, advertId uint) (application model.Application, err error) {
	err = db.Find(&application, applicationId).Error
	if err != nil {
		return model.Application{}, err
	}
	if application.UserID != userId {
		return model.Application{}, errors.New("user cannot confirm this application")
	}
	if application.AdvertID != advertId {
		return model.Application{}, errors.New("no such application for event")
	}
	application.Verified = true
	err = db.Save(application).Error
	if err != nil {
		return model.Application{}, err
	}
	return application, nil
}
