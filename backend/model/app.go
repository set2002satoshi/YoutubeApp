package model

import (
	"gorm.io/gorm"
)

type MyList struct {
	gorm.Model
	UsersID         uint   `gorm:"users_id"`
	ChannelID       string `json:"channel_id" gorm:"unique;not null"`
	ChannelName     string `json:"channel_name" gorm:"not null"`
	ChannelImageUrl string `json:"channel_image_url" gorm:"unique;not null"`
}

type AddMyListForm struct {
	ChannelID       string `json:"channel_id" gorm:"unique;not null"`
	ChannelName     string `json:"channel_name" gorm:"unique;not null"`
	ChannelImageUrl string `json:"channel_image_url" gorm:"unique;not null"`
}


