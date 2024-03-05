package service

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/e"
	"ginmall/serializer"
	"mime/multipart"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	ImgPath       string `json:"img_path" form:"img_path"`
	ProductName   string `json:"product_name" form:"product_name"`
	Price         string `json:"price" form:"price"`
	CategoryId    int    `json:"category_id" form:"category_id"`
	Title         int    `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	Onsale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.Base
}

func (service ProductService) Create(ctx context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	var err error
	var boss *model.User

	code := e.Success

	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uid)

	// 上传图片 一第一张为封面
	tmp, _ := files[0].Open()
	path, err := UploadProductToLoacalStatic(tmp, uid, service.ProductName)
}
