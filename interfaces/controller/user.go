package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai-ddd/application/user"
	logger "github.com/kangyueyue/go-ai-ddd/infrastructure/common/log"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types"
)

var ctx = context.Background()

func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(types.RegisterRequest)
		resp := new(types.RegisterResponse)

		// 接受参数
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.JSON(http.StatusOK, resp.CodeOf(types.CodeInvalidParams))
			return
		}
		token, c := user.GlobalUserServiceImpl.Register(ctx, req.Email, req.Password, req.Captcha)
		if c != types.CodeSuccess {
			ctx.JSON(http.StatusOK, resp.CodeOf(c))
			return
		}
		resp.Success()
		resp.Token = token
		logger.Log.Infof("register success,email:%s", req.Email)
		ctx.JSON(http.StatusOK, resp)
	}
}

// Login 登录
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// req := new(types.LoginRequest)
		// resp := new(types.LoginResponse)

		// // 接受参数
		// if err := ctx.ShouldBindJSON(req); err != nil {
		// 	ctx.JSON(http.StatusOK, resp.CodeOf(types.CodeInvalidParams))
		// 	return
		// }
		// token, c := user.Login(req.Username, req.Password)
		// if c != types.CodeSuccess {
		// 	ctx.JSON(http.StatusOK, resp.CodeOf(c))
		// 	return
		// }
		// resp.Success()
		// resp.Token = token
		// logger.Log.Infof("login success,username:%s", req.Username)
		// ctx.JSON(http.StatusOK, resp)
	}
}

// Captcha 验证码
func Captcha() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// req := new(types.CaptchaRequest)
		// resp := new(types.CaptchaResponse)

		// // 接受参数
		// if err := ctx.ShouldBindJSON(req); err != nil {
		// 	ctx.JSON(http.StatusOK, resp.CodeOf(types.CodeInvalidParams))
		// 	return
		// }
		// c := user.SendCaptcha(req.Email)
		// if c != types.CodeSuccess {
		// 	ctx.JSON(http.StatusOK, resp.CodeOf(c))
		// 	return
		// }
		// resp.Success()
		// logger.Log.Infof("send to email:%s success", req.Email)
		// ctx.JSON(http.StatusOK, resp)
	}
}
