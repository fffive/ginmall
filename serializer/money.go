package serializer

import (
	"ginmall/model"
	"ginmall/pkg/utils"
)

type MoneyVo struct {
	ID       uint   `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	Money    string `json:"money" form:"money"`
}

func BuildMoney(user *model.User, key string) *MoneyVo {
	utils.Encrypt.SetKey(key)
	return &MoneyVo{
		ID: user.ID,
		UserName: user.UserName,
		Money: utils.Encrypt.AesDecoding(user.Money),
	}
}
