package v1

import (
	"ginmall/pkg/utils"
	"ginmall/service"
	logging "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	var createProductService service.ProductService
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claims.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorRespond(err))
		logging.Info(err)
	}
}
