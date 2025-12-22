package user

import (
	"context"
	"sync"

	user "github.com/kangyueyue/go-ai-ddd/domain/user/service"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types"
)

// IUserService 用户服务接口
type IUserService interface {
	Register(ctx context.Context, email, passwd, captcha string) (string, types.Code)
	Login()
	Captcha()
}

// UserService 用户服务实现
type UserService struct {
	UserDomain user.IUserDomain
}

func (u *UserService) Register(ctx context.Context, email, passwd, captcha string) (string, types.Code) {
	return u.UserDomain.Register(ctx, email, passwd, captcha)
}

func (u *UserService) Login() {

}

func (u *UserService) Captcha() {

}

var (
	GlobalUserServiceImpl *UserService
	UserServiceImplOnce   sync.Once
)

// GetUserServiceImpl 获取用户服务实例
func GetUserServiceImpl(u user.IUserDomain) *UserService {
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
