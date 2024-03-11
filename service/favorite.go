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

type FavoriteService struct {
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	ProductId  uint `json:"product_id" form:"product_id"`
	model.Base
}

func (service *FavoriteService) Create(ctx context.Context, uid uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)

	exist, err := favoriteDao.FavoriteExistOrNot(service.ProductId, uid)
	if exist {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	bossDao := dao.NewUserDaoByDb(userDao.DB)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favorite := &model.Favorite{
		UserId:    uid,
		User:      *user,
		BossId:    service.BossId,
		Boss:      *boss,
		ProductId: service.ProductId,
		Product:   *product,
	}

	favoriteDao = dao.NewFavoriteDaoByDB(favoriteDao.DB)
	err = favoriteDao.Create(favorite)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) Show(ctx context.Context, uid uint) serializer.Response {
	favoriteDao := dao.NewFavoriteDao(ctx)
	var code int

	// 展示判断 pagesize
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	favorites, err := favoriteDao.Show(uid, service.Base)
	if err != nil {
		logrus.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(len(favorites)))
}

func (service *FavoriteService) Delete(ctx context.Context, uid uint, id string) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoriteDao(ctx)
	fid, _ := strconv.Atoi(id)
	service.FavoriteId = uint(fid)

	err := favoriteDao.Delete(service.FavoriteId)
	if err != nil {
		logrus.Info(err)
		code := e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "成功取消收藏",
	}
}
