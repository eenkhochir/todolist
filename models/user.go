package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Phonenumber uint   `json:"phonenumber" validate:"required"`
	Tasks       []Todo `json:"tasks" gorm:"foreignKey:Userid"`
}
