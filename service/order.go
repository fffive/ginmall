package service

import (
	"ginmall/model"
)

type OrderService struct {
	ProductID uint `form:"product_id" json:"product_id"`
	AddressID uint `form:"address_id" json:"address_id"`
	BossID    uint `form:"boss_id" json:"boss_id"`
	UserID    uint `form:"user_id" json:"user_id"`
	Money     int  `form:"money" json:"money"`
	OrderNum  uint `form:"order_num" json:"order_num"`
	Num       uint `form:"num" json:"num"`
	Type      int  `form:"type" json:"type"`
	model.Base
}

// func (service OrderService) Create(ctx context.Context, id uint) serializer.Response {
// }
