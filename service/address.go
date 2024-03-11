package service

import (
	"context"
	"ginmall/dao"
	"ginmall/model"
	"ginmall/pkg/e"
	"ginmall/serializer"
	"strconv"

	"github.com/sirupsen/logrus"
)

type AddressService struct {
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
}

func (service *AddressService) Create(ctx context.Context, uid uint) serializer.Response {
	var code int
	addressDao := dao.NewAddressDao(ctx)
	address := &model.Address{
		UserId:  uint(uid),
		Name:    service.Name,
		Address: service.Address,
		Phone:   service.Phone,
	}

	err := addressDao.Create(address)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao = dao.NewAddressDaoByDB(addressDao.DB)
	addresses, err := addressDao.GetAddressByUid(uid)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildAddresses(addresses), uint(len(addresses)))
}

// 查询单个收藏夹
func (service *AddressService) List(ctx context.Context, uid uint, aid string) serializer.Response {
	var code int
	addressDao := dao.NewAddressDao(ctx)

	address, err := addressDao.List(uid, aid)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildAddress(address),
		Msg:    e.GetMsg(code),
	}
}

// 删除地址
func (service *AddressService) Delete(ctx context.Context, uid uint, aid string) serializer.Response {
	var code int
	addressDao := dao.NewAddressDao(ctx)

	err := addressDao.Delete(uid, aid)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   "成功删除地址",
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Show(ctx context.Context, uid uint) serializer.Response {
	var code int
	addressDao := dao.NewAddressDao(ctx)

	address, err := addressDao.ListAddressById(uid)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   serializer.BuildAddresses(address),
		Msg:    e.GetMsg(code),
	}
}

func (service *AddressService) Update(ctx context.Context, uid uint, aid string) serializer.Response {
	var code int
	addressDao := dao.NewAddressDao(ctx)
	aId, _ := strconv.Atoi(aid)

	address := &model.Address{
		UserId: uid,
		Name:   service.Name,
		Phone:  service.Phone,
		Address: service.Address,
	}

	err := addressDao.UpdateAddress(uint(aId), address)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	addressDao = dao.NewAddressDaoByDB(addressDao.DB)
	var addresses []*model.Address
	addresses, _ = addressDao.GetAddressByUid(uid)

	return serializer.Response{
		Status: code,
		Data:  	serializer.BuildAddresses(addresses),
		Msg:    e.GetMsg(code),
	}
}
