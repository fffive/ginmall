package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) AddressDao {
	return AddressDao{NewDBclient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) AddressDao {
	return AddressDao{db}
}

// 创建address
func (dao AddressDao) Create(address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&address).Error
}

// 根据uid获取地址
func (dao AddressDao)GetAddressByUid(uid uint) (address []*model.Address, err error) {
	err = dao.DB.Where("user_id = ?", uid).Find(&address).Error
	return
}

//查询单个
func (dao AddressDao) List(uid uint, aid string) (address *model.Address, err error) {
	err = dao.DB.Where("user_id = ? AND id = ?",uid, aid).First(&address).Error

	return 
}

func (dao AddressDao) Delete(uid uint, aid string) error {
	return dao.DB.Where("user_id = ? AND id = ?", uid, aid).Delete(&model.Address{}).Error
}
