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

// 展示所有
func ListProducts(c *gin.Context) {
	ListProductsService := service.ProductService{}

	if err := c.ShouldBind(&ListProductsService); err == nil {
		res := ListProductsService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Info(err)
	}
}

// 查询
func SearchProducts(c *gin.Context) {
	searchProductsService := service.ProductService{}

	if err := c.ShouldBind(&searchProductsService); err == nil {
		res := searchProductsService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		logging.Info(err)
	}
}

// 获取单个商品详细信息
func ShowProduct(c *gin.Context) {
	showProductService := service.ProductService{}
	res := showProductService.Show(c.Request.Context(), c.Param("id"))
	c.JSON(http.StatusOK, res)
}

// 获取商品图片 给予前端展示
func ListProductImg(c *gin.Context) {
	listProductImg := service.ListProductImgService{}
	if err := c.ShouldBind(&listProductImg); err == nil {
		res := listProductImg.ListImg(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	}else {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, err)
	}
}