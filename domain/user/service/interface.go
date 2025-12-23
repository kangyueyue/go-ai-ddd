package user

import (
	"context"

	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
)

// IUserDomain 接口
type IUserDomain interface {
	Login(ctx context.Context, email, passwd string) (string, code.Code)
	Register(ctx context.Context, email, passwd, captcha string) (string, code.Code)
	Captcha(ctx context.Context, email string) code.Code
}
