package serializer

import (
	"ginmall/model"
)

type AddressVo struct {
	ID       uint   `json:"id" form:"id"`
	UserId   uint   `json:"user_id" form:"user_id"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Seen     bool   `json:"seen" form:"seen"`
	CreateAt int64  `json:"create_at" form:"create_at"`
}

func BuildAddress(item *model.Address) AddressVo {
	return AddressVo{
		ID: item.ID,
		UserId: item.UserId,
		Name: item.Name,
		Phone: item.Phone,
		Address: item.Address,
		Seen: false,
		CreateAt: item.CreatedAt.Unix(),
	}
}

func BuildAddresses(items []*model.Address) (addresses []AddressVo) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return
}
