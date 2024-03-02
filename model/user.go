package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Email    string
	PassWord string
	NickName string
	Status   string
	Avatar   string
	Money    string
}
