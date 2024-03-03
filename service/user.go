package service

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/e"
	"ginmall/pkg/utils"
	"ginmall/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"pass_word" form:"pass_word"`
	Key      string `json:"key" form:"key"`
}

func (service UserService) Register(ctx context.Context) serializer.Response {
	var user model.User

	code := e.Success
	if service.Key == "" || len(service.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "秘钥长度不足",
		}
	}

	// 加密算法
	utils.Encrypt.SetKey(service.Key)

	// dao *Dao := dao.*(ctx)
	userDao := dao.NewUserDao(ctx)

	// 用户验证是否存在
	_, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if err != nil {
		code = e.ErrDatabase
		return serializer.Response{
			Status: code,	
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorUserExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Avatar:   "default_avatar",
		Status:   model.Active,
		Money:    utils.Encrypt.AesEncoding("10000"),
	}

	// 用户密码加密
	if err = user.SetPassWord(service.PassWord); err != nil {
		code = e.ErrFailEncrypte
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	// 用户创建
	if err = userDao.CreateUser(&user); err != nil {
		code = e.Error
	}
	return serializer.Response{
		// Status: http.StatusBadRequest,
		Status: code,
		Msg: e.GetMsg(code),
	}
}
