package service

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/e"
	"ginmall/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	ImgPath       string `json:"img_path" form:"img_path"`
	ProductName   string `json:"product_name" form:"product_name"`
	Price         string `json:"price" form:"price"`
	CategoryId    int    `json:"category_id" form:"category_id"`
	Title         string    `json:"title" form:"title"`
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
	if err != nil {
		code = e.ErrorImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	product := &model.Product{
		ProductName:   service.ProductName,
		Price:         service.Price,
		DiscountPrice: service.ProductName,
		Info:          service.Info,
		Title:         service.Title,
		CategoryId:    service.CategoryId,
		Onsale:        true,
		Num:           service.Num,
		ImgPath:       path,
		BossId:        uid,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	productDao := dao.NewProductDao(ctx)
	err = productDao.Create(product)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 并发上传图片
	wg := new(sync.WaitGroup)
	wg.Add(len(files))

	for index, file := range files {
		num := strconv.Itoa(index)
		ProductImgDao := dao.NewProductImgDaoByDb(productDao.DB)
		tmp, _ := file.Open()

		path, err := UploadProductToLoacalStatic(tmp, uid, service.ProductName+num)
		if err != nil {
			code = e.ErrorImgUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := &model.ProductImg{
			ProductImgId: product.ID,
			ImgPath:      path,
		}

		err = ProductImgDao.Create(productImg)
		if err != nil {
			code = e.ErrorDatabase
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}

		wg.Done()
	}
	
	wg.Wait()
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}
