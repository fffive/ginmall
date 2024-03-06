package serializer

import (
	"ginmall/conf"
	"ginmall/model"
)

type ProductVo struct {
	Id            uint   `json:"id" form:"id"`
	ProductImg    string `json:"product_img" form:"product_img"`
	ProductName   string `json:"product_name" form:"product_name"`
	Price         string `json:"price" form:"price"`
	ImgPath       string `json:"img_path" form:"img_path"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	CategoryId    int    `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	View          uint64 `json:"view" form:"view"`
	CreateAt      int64  `json:"create_at" form:"create_at"`
	Onsale        bool   `json:"on_sale" form:"on_sale"`
	Info          string `json:"info" form:"info"`
	Num           int    `json:"num" form:"num"`
	BossId        uint   `json:"boss_id" form:"boss_id"`
	BossName      string `json:"boss_name" form:"boss_name"`
	BossAvatar    string `json:"boss_avatar" form:"boss_avatar"`
}

func BuildProduct(product *model.Product) *ProductVo {
	return &ProductVo{
		Id:            product.ID,
		Onsale:        product.Onsale,
		ProductImg:    product.ImgPath,
		ProductName:   product.ProductName,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		CategoryId:    product.CategoryId,
		View:          product.View(),
		CreateAt:      product.CreatedAt.Unix(),
		Title:         product.Title,
		Info:          product.Info,
		Num:           product.Num,
		BossId:        product.BossId,
		BossName:      product.BossName,
		BossAvatar:    product.BossAvatar,
	}
}
