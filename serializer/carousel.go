package serializer

import "ginmall/model"

type CarouselVo struct {
	Id uint `json:"id"`
	ImgePath string `json:"img_path"`
	ProductId uint `json:"product_id"`
	CreateAt int64 `json:"create_at"` 
}

func BulidCarousel(Item *model.Carousel) CarouselVo {
	return CarouselVo {
		Id: Item.ID,
		ImgePath: Item.ImgPath,
		ProductId: Item.ProductId,
		CreateAt: Item.CreatedAt.Unix(),
	}
}

func BuildCarousels(items []model.Carousel) (carousels []CarouselVo) {
	for _, item := range items {
		carousel := BulidCarousel(&item)
		carousels = append(carousels, carousel)
	}
	return 
}