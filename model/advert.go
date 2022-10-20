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
	UserID      uint      `json:"user_id" validate:"required"`
}

type CreateAdvertDTO struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	City        string    `json:"city" validate:"required"`
	Date        time.Time `json:"date" validate:"required"`
	UserID      uint      `json:"user_id" validate:"required"`
}
