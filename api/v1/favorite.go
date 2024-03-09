package v1

import (
	"ginmall/pkg/utils"
	"ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 创建
func CraeteFavorite(c *gin.Context) {
	service := service.FavoriteService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}

// 展示
func ShowFavorities(c *gin.Context) {
	service := service.FavoriteService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}

// 删除
func DeleteFavorite(c *gin.Context) {
	service := service.FavoriteService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context(), claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}
