package dao

import (
	"context"

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

