package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	AdminName     string `gorm:"type:varchar(15) not null"`
	AdminPassWord string `gorm:"type:varchar(15) not null"`
	Avatar        string
}
