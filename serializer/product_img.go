package serializer

import (
	"ginmall/conf"
	"ginmall/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ID,
		ImgPath: conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		productImg := BuildProductImg(item)
		productImgs = append(productImgs, productImg)
	}
	return 
}

