package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai-ddd/application/user"
	logger "github.com/kangyueyue/go-ai-ddd/infrastructure/common/log"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/dto"
)

// ctx
var ctx = context.Background()

// Register 注册
func Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(dto.RegisterRequest)
		resp := new(dto.RegisterResponse)

		// 接受参数
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.JSON(http.StatusOK, resp.CodeOf(code.CodeInvalidParams))
			return
		}
		token, c := user.GlobalUserServiceImpl.Register(ctx, req.Email, req.Password, req.Captcha)
		if c != code.CodeSuccess {
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
		req := new(dto.LoginRequest)
		resp := new(dto.LoginResponse)

		// 接受参数
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.JSON(http.StatusOK, resp.CodeOf(code.CodeInvalidParams))
			return
		}
		token, c := user.GlobalUserServiceImpl.Login(ctx, req.Username, req.Password)
		if c != code.CodeSuccess {
			ctx.JSON(http.StatusOK, resp.CodeOf(c))
			return
		}
		resp.Success()
		resp.Token = token
		logger.Log.Infof("login success,username:%s", req.Username)
		ctx.JSON(http.StatusOK, resp)
	}
}

// Captcha 验证码
func Captcha() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(dto.CaptchaRequest)
		resp := new(dto.CaptchaResponse)

		// 接受参数
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.JSON(http.StatusOK, resp.CodeOf(code.CodeInvalidParams))
			return
		}
		c := user.GlobalUserServiceImpl.Captcha(ctx, req.Email)
		if c != code.CodeSuccess {
			ctx.JSON(http.StatusOK, resp.CodeOf(c))
			return
		}
		resp.Success()
		logger.Log.Infof("send to email:%s success", req.Email)
		ctx.JSON(http.StatusOK, resp)
	}
}
