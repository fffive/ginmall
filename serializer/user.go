package serializer

import (
	"ginmall/conf"
	"ginmall/model"
)

// 序列化成前端需要额user
type UserVo struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Status   string `json:"status"`
	Email    string `json:"email"`
	Type     int    `json:"type"`
	CreateAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *UserVo {
	return &UserVo{
		ID: user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Avatar: conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		Status: user.Status,
		Email: user.Status,
		CreateAt: user.CreatedAt.Unix(),
	}
}