package model

import (
	"time"

	"gorm.io/gorm"
)

type Advert struct {
	gorm.Model
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
}

type Application struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	AdvertID uint `json:"advert_id"`
	Verified bool `json:"verified"`
}

type CreateAdvertDTO struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}

type UpdateAdvertDTO struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
}
