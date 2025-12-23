package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai-ddd/interfaces/controller"
)

// UserRouter
func UserRouter(v1 *gin.RouterGroup) {
	v1.POST("/register", controller.Register())
	v1.POST("/login", controller.Login())
	v1.POST("/captcha", controller.Captcha())
}
