package routes

import (
	api "ginmall/api/v1"
	"ginmall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
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
		v1.POST("user/login", api.UserLogin)

		// 轮播图的展示
		v1.GET("carousels", api.ListCarousels)

		//商品操作
		v1.GET("products", api.ListProducts)
		v1.GET("product/:id", api.ShowProduct)
		v1.POST("products", api.SearchProducts)
		v1.GET("imgs/:id", api.ListProductImg)
		v1.GET("categories", api.ListCategories)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			// 修改用户信息
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			// 查看Money
			authed.POST("money", api.ShowMoney)

			// 用户创建商品
			authed.POST("product", api.CreateProduct)

			//收藏夹
			authed.POST("favorites", api.CraeteFavorite)
			authed.GET("favorites", api.ShowFavorities)
			authed.DELETE("favorites/:id", api.DeleteFavorite)
		}
	}

	return r
}
