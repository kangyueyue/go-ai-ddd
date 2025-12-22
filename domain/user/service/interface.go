package user

import (
	"context"

	"github.com/kangyueyue/go-ai-ddd/interfaces/types"
)

// IUserDomain 接口
type IUserDomain interface {
	Login(ctx context.Context, email, passwd, captcha string)
	Register(ctx context.Context, email, passwd, captcha string) (string, types.Code)
	Captcha(ctx context.Context)
}
