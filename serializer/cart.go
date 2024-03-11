package serializer

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
)

type CartVo struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check         bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	Desc          string `json:"desc"`
}

func BulidCart(cart *model.Cart, product *model.Product, boss *model.User) *CartVo {
	return &CartVo{
		ID:            cart.ID,
		UserID:        cart.UserId,
		ProductID:     product.ID,
		CreateAt:      cart.CreatedAt.Unix(),
		Num:           cart.Num,
		MaxNum:        cart.MaxNum,
		Check:         cart.Check,
		Name:          product.ProductName,
		ImgPath:       product.ImgPath,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossName:      boss.UserName,
		Desc:          product.Info,
	}
}

func BuildCarts(items []*model.Cart) (carts []*CartVo) {
	for _, item := range items {
		product, err := dao.NewProductDao(context.Background()).GetProductById(item.Product)
		if err != nil {
			continue
		}

		boss, err := dao.NewUserDao(context.Background()).GetUserById(item.BossId)
		if err != nil {
			continue
		}

		cart := BulidCart(item, product, boss)
		carts = append(carts, cart)
	}
	return
}
