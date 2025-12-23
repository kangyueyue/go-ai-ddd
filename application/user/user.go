package user

import (
	"context"
	"sync"

	user "github.com/kangyueyue/go-ai-ddd/domain/user/service"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
)

// IUserService 用户服务接口
type IUserService interface {
	Register(ctx context.Context, email, passwd, captcha string) (string, code.Code)
	Login(ctx context.Context, username, passwd string) (string, code.Code)
	Captcha(ctx context.Context, email string) code.Code
}

// UserService 用户服务实现
type UserService struct {
	UserDomain user.IUserDomain
}

func (u *UserService) Register(ctx context.Context, email, passwd, captcha string) (string, code.Code) {
	return u.UserDomain.Register(ctx, email, passwd, captcha)
}

// Login 登入
func (u *UserService) Login(ctx context.Context, username, passwd string) (string, code.Code) {
	return u.UserDomain.Login(ctx, username, passwd)
}

func (u *UserService) Captcha(ctx context.Context, email string) code.Code {
	return u.UserDomain.Captcha(ctx, email)
}

var (
	GlobalUserServiceImpl *UserService
	UserServiceImplOnce   sync.Once
)

// GetUserServiceImpl 获取用户服务实例
func GetUserServiceImpl(u user.IUserDomain) IUserService {
	if GlobalUserServiceImpl != nil {
		return GlobalUserServiceImpl
	}
	UserServiceImplOnce.Do(func() {
		if GlobalUserServiceImpl == nil {
			GlobalUserServiceImpl = &UserService{
				UserDomain: u,
			}
		}
	})
	return GlobalUserServiceImpl
}
