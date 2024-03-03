package dao

import (
	"context"
	"ginmall/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// 请求上下文
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBclient(ctx)}
}

func NewUserDaoByDb(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// username 检查是否存在 --- 用户名不能一致
func (dao *UserDao) ExistOrNotByUserName(username string) (user *model.User, exist bool, err error) {
	var count int64

	err = dao.DB.Model(&model.User{}).Where("user_name=?", username).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	err = dao.DB.Model(&model.User{}).Where("user_name=?", username).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// 创建用户
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) UpdateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Updates(&user).Error
}

func (dao *UserDao) GetUserById(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
}
