package routes

import (
	"ginmall/middleware"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	r := gin.Default()

	r.Use(middleware.Cors())
	// 加载静态文件
	r.StaticFS("/static", http.Dir(".static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "success")
		})
		// 用户操作 
		v1.POST("user/register", api.UserRegister)
		
	}
	
	return r
}