package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai-ddd/interfaces/adapter/user"
)

// NewRouter 初始化路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	group := r.Group("api/v1/")
	{
		// 服务连通性测试
		group.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		// 用户相关接口
		user.UserRouter(group.Group("/user"))
		// TODO: session对话接口，需要加鉴权
	}
	return r
}
