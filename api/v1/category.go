package v1

import (
	"ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ListCategories(c *gin.Context) {
	listCategoryService := service.ListCategoryService{}

	if err := c.ShouldBind(&listCategoryService); err == nil {
		res := listCategoryService.ListCategory(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else {
		logrus.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}
