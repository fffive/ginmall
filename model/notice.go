package model

import "gorm.io/gorm"

type Notice struct {
	gorm.Model
	Info string `gorm:"type:text"`
}
