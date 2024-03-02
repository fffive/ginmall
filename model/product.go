package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ImgPath       string
	ProductName   string
	Price         string
	Category      int
	Title         int
	Info          string
	DiscountPrice string
	Onsale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}
