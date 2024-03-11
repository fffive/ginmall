package dao

import (
	"context"
	"ginmall/model"
	"ginmall/pkg/e"

	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBclient(ctx)}
}

func NewCartDaoByDB(db CartDao) *CartDao {
	return &CartDao{db.DB}
}

func (dao *CartDao) CreateCart(pid, uid, bid uint) (cart *model.Cart, status uint, err error) {
	cart, err = dao.GetCartById(pid, uid, bid)

	// 是否获取到 没有就创建 有就增加 超过就报错
	if err == gorm.ErrRecordNotFound {
		cart = &model.Cart{
			Product: pid,
			UserId:  uid,
			BossId:  bid,
			Num:     1,
			MaxNum:  10,
			Check:   false,
		}
		err = dao.DB.Create(&cart).Error
		if err != nil {
			return
		}
		return cart, e.Success, err
	} else if cart.Num < cart.MaxNum {
		cart.Num++
		err = dao.DB.Save(&cart).Error
		if err != nil {
			return
		}
		return cart, e.ErrorProductExistCart, err
	} else {
		return cart, e.ErrorProductMoreCart, err
	}
}

func (dao CartDao) GetCartById(pid, uid, bid uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).
		Where("product = ? AND user_id = ? AND boss_id = ?", pid, uid, bid).
		First(&cart).Error

	return
}
