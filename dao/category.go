package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBclient(ctx)}
}

func NewCategoryDaoByDb(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

func (dao *CategoryDao) List() (categories []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&categories).Error
	return
}
