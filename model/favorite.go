package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeighKey:UserId not null"`
	UserId    uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeighKey:UserId not null"`
	BossId    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeighKey:ProductId not null"`
	ProductId uint    `gorm:"not null"`
}
