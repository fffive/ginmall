package v1

import (
	"ginmall/pkg/utils"
	"ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 收获地址的创建
func CreateAddress(c *gin.Context) {
	service := service.AddressService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(),  claims.ID)
		c.JSON(http.StatusOK, res)
	}else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}

// 获取单个地址信息 
func GetAddress(c *gin.Context) {
	service := service.AddressService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c.Request.Context(),  claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	}else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}

// 删除地址
func DeleteAddress(c *gin.Context) {
	service := service.AddressService{}

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context(),  claims.ID, c.Param("id"))
		c.JSON(http.StatusOK, res)
	}else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}