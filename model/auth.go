package model

type UserClaims struct {
	ID    uint
	Email string
}

type AuthRegisterDTO struct {
	Username string `json:"username" validate:"required" gorm:"unique"`
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

type AuthLoginDTO struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}
