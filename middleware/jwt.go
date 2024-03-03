package middleware

import (
	"ginmall/pkg/e"
	"ginmall/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200

		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		}else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			}else if time.Now().Unix() > claims.ExpiresAt{
				code = e.ErrorAuthCheckTokenTimeOut
			}
		}
		if code != e.Success {
			c.JSON(http.StatusOK, gin.H {
				"status": code,
				"msg": e.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}