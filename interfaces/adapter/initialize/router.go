package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai-ddd/interfaces/controller"
)

// NewRouter 初始化路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	// TODO:跨域中间件
	v1 := r.Group("api/v1/")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户操作
		v1.POST("register", controller.Register())
		v1.POST("login", controller.Login())
		v1.POST("captcha", controller.Captcha())
		{
			// TODO: session，需要加鉴权
		}
	}
	return r
}
