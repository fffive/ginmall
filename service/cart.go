package service

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/e"
	"ginmall/serializer"

	"github.com/sirupsen/logrus"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	Bossid    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       uint `json:"num" form:"num"`
}

func (service CartService) Create(ctx context.Context, id uint) serializer.Response {
	var code int
	var product *model.Product

	// 判断上平是否存在 并获取商品信息
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 获取是否存在购物车 并创建购物车和增加购物车中商品数量
	cartDao := dao.NewCartDao(ctx)
	cart, status, err := cartDao.CreateCart(service.ProductId, id, service.Bossid)
	if status == e.ErrorProductMoreCart {
		logrus.Info(err)
		return serializer.Response{
			Status: int(status),
			Msg:    e.GetMsg(int(status)),
		}
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(id)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: int(status),
		Msg: e.GetMsg(int(status)),
		Data: serializer.BulidCart(cart, product, boss),
	}
}
