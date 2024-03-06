package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBclient(ctx)}
}

func NewProductDaoByDb(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

func (dao ProductDao) Create(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}
