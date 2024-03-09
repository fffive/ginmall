package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBclient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao FavoriteDao) Create(favorite *model.Favorite) (err error) {
	err = dao.DB.Create(&favorite).Error
	return
}

func (dao FavoriteDao) FavoriteExistOrNot(pid, uid uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).
		Where("user_id = ? AND product_id = ?", uid, pid).Count(&count).Error
	if count == 0 || err != nil {
		return false, err
	}
	return true, err
}

func (dao FavoriteDao) Delete(fid uint) error {
	return dao.DB.Where("id = ?", fid).Delete(&model.Favorite{}).Error
}

func (dao FavoriteDao) Show(uid uint, page model.Base) (favorites []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Preload("User").
		Where("user_id = ?", uid).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&favorites).Error
	return 
}
