package v1

import (
	"ginmall/pkg/utils"
	"ginmall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	createCartService := service.CartService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
		utils.LogrusObj.Infoln(err)
	}
}
