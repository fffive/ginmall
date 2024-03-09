package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao {
	return &ProductImgDao{NewDBclient(ctx)}
}

func NewProductImgDaoByDb(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao ProductImgDao) Create(ProductImg *model.ProductImg) error {
	return dao.DB.Model(&model.ProductImg{}).Create(&ProductImg).Error
}

func (dao ProductImgDao) List(pid uint) (productImg []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).
		Where("product_img_id = ?", pid).
		Find(&productImg).Error
	return
}
