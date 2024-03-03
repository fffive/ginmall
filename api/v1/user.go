package v1

import (
	"net/http"

	"ginmall/service"
	"github.com/gin-gonic/gin"
)

// 编写 Http 请求 使用*gin.Context
// 注册接口
func UserRegister(c *gin.Context) {
	// 一个 service 类
	var userRegister service.UserService
	
	// 进行绑定判断 返回respond json respond
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// 登录接口 http请求
func UserLogin(c *gin.Context) {
	var userLogin service.UserService

	// 进行绑定判断 返回respond json 
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, err)
	}
}