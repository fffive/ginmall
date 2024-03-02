package service

import (
	"context"
	"ginmall/model"
	"ginmall/pkg/e"
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
		
	}
}