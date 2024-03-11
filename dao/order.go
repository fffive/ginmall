package dao

import (
	"context"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) OrderDao {
	return OrderDao{NewDBclient(ctx)}
}

func NewOrderDaoByDB(db *gorm.DB) OrderDao {
	return OrderDao{db}
}


