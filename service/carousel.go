package service

import (
	"context"

	"ginmall/dao"
	"ginmall/pkg/e"
	"ginmall/serializer"
	logging "github.com/sirupsen/logrus"
)

type ListCarouselsService struct {
}

func (service ListCarouselsService) List(ctx context.Context) serializer.Response {
	code := e.Success
	carouselDao := dao.NewCarouselDao(ctx)

	carousel, err := carouselDao.ListAddress()
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousel), uint(len(carousel)))
}
