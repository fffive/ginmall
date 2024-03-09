package v1

import (
	"ginmall/pkg/utils"
	"ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CraeteFavorite(c *gin.Context) {
	service := service.FavoriteService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}

func ShowFavorities(c *gin.Context) {
	service := service.FavoriteService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}


