package service

import (
	"context"
	"ginmall/dao"
	"ginmall/pkg/e"
	"ginmall/serializer"
)

type ListCarouselsService struct {
}

func (service ListCarouselsService) List(ctx context.Context) serializer.Response {
	code := e.Success
	carouselDao := dao.NewCarouselDao(ctx)

	carousel, err := carouselDao.ListAddress()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousel),uint(len(carousel)))
}