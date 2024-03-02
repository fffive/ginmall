package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService

	if err := c.ShouldBindJSON(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context()) // 传递上下文
		c.JSON(http.StatusOK, res)
	}
}
