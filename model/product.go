package model

import (
	"ginmall/cache"
	"strconv"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ImgPath       string
	ProductName   string
	Price         string
	CategoryId    int
	Title         string
	Info          string
	DiscountPrice string
	Onsale        bool `gorm:"default:false"`
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

// 实现点击率的查看
func (Product *Product) View() uint64 {
	cache.InitCache()
	// 字符串转换直接显示结果
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(Product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)

	return count
}

// 商品浏览
func (Product *Product) AddView() {
	// 点击数增加
	cache.RedisClient.Incr(cache.ProductViewKey(Product.ID))
	// 增加排行点击率
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Product.ID)))
}
