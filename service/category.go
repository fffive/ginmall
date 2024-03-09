package service

import (
	"context"
	"ginmall/dao"
	"ginmall/pkg/e"
	"ginmall/serializer"

	"github.com/sirupsen/logrus"
)

type ListCategoryService struct {
}

func (service ListCategoryService) ListCategory(ctx context.Context) serializer.Response {
	code := e.Success
	categoryDao := dao.NewCategoryDao(ctx)

	categories, err := categoryDao.List()
	if err != nil {
		logrus.Info(err)
		code = e.ErrorCategoryGetFail
		return serializer.Response{
			Status: code,
			Msg: e.GetMsg(code),
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data: serializer.BuildCategories(categories),
	}
}
