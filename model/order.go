package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	UserId    uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num       int
	OrderNum  uint64
	Type      uint // 1 未支付 2 已支付
	Money     float64
}
