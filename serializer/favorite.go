package serializer

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
)

type FavoriteVo struct {
	UserId        uint   `json:"user_id" form:"user_id"`
	BossId        uint   `json:"boss_id" form:"boss_id"`
	ProductId     uint   `json:"product_id" form:"product_id"`
	CreateAt      int64  `json:"create_at" form:"create_at"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPricr string `json:"discount_price" form:"discount_price"`
	Num           int    `json:"num" form:"num"`
	Onsale        bool   `json:"on_sale" form:"on_sale"`
}

func BuildFavorite(item1 *model.Favorite, item2 *model.Product, item3 *model.User) FavoriteVo {
	return FavoriteVo{
		UserId: item1.UserId,
		BossId: item3.ID,
		ProductId: item2.ID,
		CreateAt: item1.CreatedAt.Unix(),
		Name: item2.ProductName,
		CategoryId: uint(item2.CategoryId),
		Title: item2.Title,
		Info: item2.Info,
		ImgPath: item2.ImgPath,
		Price: item2.Price,
		DiscountPricr: item2.DiscountPrice,
		Num: item2.Num,
		Onsale: item2.Onsale,
	}
}

func BuildFavorites(ctx context.Context, items []*model.Favorite) (favorites []FavoriteVo) {
	bossDao := dao.NewUserDao(ctx)
	ProductDao := dao.NewProductDao(ctx)

	for _, item := range items {
		boss, _ := bossDao.GetUserById(item.UserId)
		product, _ := ProductDao.GetProuctById(item.ProductId)

		favorite := BuildFavorite(item, product, boss)
		favorites = append(favorites, favorite)
	}
	return 
}