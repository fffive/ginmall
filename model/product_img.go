package model

import "gorm.io/gorm"

type ProductImg struct {
	gorm.Model
	ProductImgId uint
	ImgPath      string
}