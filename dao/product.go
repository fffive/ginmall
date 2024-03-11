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

// 根据情况获取商品数量以展示
func (dao ProductDao) CountProductByConditon(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// 获取商品列表
func (dao ProductDao) ListProductByConditon(condition map[string]interface{}, page model.Base) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error

	return
}

func (dao ProductDao) SearchProducts(info string, page model.Base) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("product_name LIKE ? or info LIKE ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&products).Error

	return
}

func (dao ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id = ?", id).First(&product).Error
	return
}
