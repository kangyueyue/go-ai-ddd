package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	myjwt "github.com/kangyueyue/go-ai-ddd/infrastructure/utils/jwt"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/dto"
)

// Auth 认证中间件
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := new(dto.Response)
		authHeader := ctx.GetHeader("Authorization")
		var token string
		// bearer 持票人模式
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// 兼容 URL 参数传 token
			token = ctx.Query("token")
		}
		if token == "" {
			ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidToken))
			ctx.Abort() // 中断当前请求
			return
		}
		username, ok := myjwt.ParseToken(token)
		if !ok {
			ctx.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidToken))
			ctx.Abort()
			return
		}
		ctx.Set("userName", username) // 后续ctx直接取
		ctx.Next()                    // 通过
	}
}
