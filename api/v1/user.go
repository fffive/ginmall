package v1

import (
	"net/http"
	logging "github.com/sirupsen/logrus"
	"ginmall/pkg/utils"
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
		c.JSON(http.StatusBadRequest, ErrorRespond(err))
		logging.Info(err)
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
		c.JSON(http.StatusBadRequest, ErrorRespond(err))
		logging.Info(err)
	}
}

// 用户信息更新
func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService

	// 校验Token
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, ErrorRespond(err))
		logging.Info(err)
	}
}

// 上传头像
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size

	var uploadAvatar service.UserService
	// 校验token
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(c.Request.Context(), claims.ID, file, fileSize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorRespond(err))
		logging.Info(err)
	}
}

// 发送email
func SendEmail(c *gin.Context) {
	var sendService service.SendService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendService); err == nil {
		res := sendService.Send(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, ErrorRespond(err))
		logging.Info(err)
	}
}

// 验证邮箱
func ValidEmail(c *gin.Context) {
	var validEmailService service.ValidEmailService

	if err := c.ShouldBind(&validEmailService); err == nil {
		res := validEmailService.Valid(c.Request.Context(), c.GetHeader("Authorization"))
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, err)
	}
}

// 展示金额 来源:中国银行的设置
func ShowMoney(c *gin.Context) {
	var showMoneyService service.ShowMoneyService

	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err == nil {
		res := showMoneyService.ShowMoney(c.Request.Context(), claims.ID)
		c.JSON(http.StatusOK, res)
	}else {
		c.JSON(http.StatusBadRequest, ErrorRespond(err))
		logging.Info(err)
	}
}

