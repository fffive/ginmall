package service

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/e"
	"ginmall/pkg/utils"
	"ginmall/serializer"
	"mime/multipart"
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
		code = e.ErrorDatabase
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

	// 不存在就可以创建用户
	user = model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Avatar:   "avatar.jpg",
		Status:   model.Active,
		Money:    utils.Encrypt.AesEncoding("10000"),
	}

	// 用户密码加密
	if err = user.SetPassWord(service.PassWord); err != nil {
		code = e.ErrorFailEncrypte
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
		Msg:    e.GetMsg(code),
	}
}

func (service UserService) Login(ctx context.Context) serializer.Response {
	var user *model.User

	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(service.UserName)
	if !exist || err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在，请先注册",
		}
	}

	// 检查密码是否正确
	if !user.CheckPassword(service.PassWord) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "密码错误，请重试",
		}
	}

	// token 签发售
	token, err := utils.GenerateToken(user.ID, user.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
	}

}

// 用户信息更新
func (service UserService) Update(ctx context.Context, uid uint) serializer.Response {
	var user *model.User
	var err error 
	code := e.Success
	userDao := dao.NewUserDao(ctx)

	// 寻找用户
	user, _ = userDao.GetUserById(uid);

	// 更改NickName
	if service.NickName != "" {
		user.NickName = service.NickName
	}

	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: "用户昵称修改已修改为"+service.NickName,
	}
}

// 修改头像
func (service UserService) Post(ctx context.Context, uid uint, file multipart.File, filesize int64) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uid)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	// 保存图片到本地
	path, err := UploadAvatarToLoacalStatic(file, uid, user.UserName)
	if err != nil {
		code = e.ErrorUploadFail
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Data: "上传文件失败",
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildUser(user),
	}
}