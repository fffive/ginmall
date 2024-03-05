package v1

import (
	"net/http"

	"ginmall/service"
	"github.com/gin-gonic/gin"
)

// 轮播图的展示
func ListCarousels(c *gin.Context) {
	var listCarouselsService service.ListCarouselsService

	if err := c.ShouldBind(&listCarouselsService); err == nil {
		res := listCarouselsService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, err)
	}
}