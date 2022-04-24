package model

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     *string `json:"user" gorm:"unique;not null"`
	Email    *string `json:"email" gorm:"unique;not null"`
	Password []byte  `json:"password" gorm:"not null"`
}

type CreateUsers struct {
	gorm.Model
	Name     string `json:"user" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

type LoginUser struct {
	Email    string `json:"email" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
