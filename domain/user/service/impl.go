package user

import (
	"context"

	"github.com/kangyueyue/go-ai-ddd/consts"
	"github.com/kangyueyue/go-ai-ddd/domain/user/entity"
	"github.com/kangyueyue/go-ai-ddd/domain/user/repository"
	myemail "github.com/kangyueyue/go-ai-ddd/infrastructure/common/email"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/redis"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/utils"
	myjwt "github.com/kangyueyue/go-ai-ddd/infrastructure/utils/jwt"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
)

// UserDomainImpl 用户领域实现
type UserDomainImpl struct {
	repo repository.IUserRepository
}

// NewUserDomain
func NewUserDomainImpl(repo repository.IUserRepository) IUserDomain {
	return &UserDomainImpl{repo: repo}
}

// Login 登入
func (u *UserDomainImpl) Login(ctx context.Context, username, password string) (string, code.Code) {
	var ok bool
	var UserInformation *entity.UserEntity
	// 1.判断用户是否存在
	if ok, UserInformation = u.repo.IsExistUserByUsername(username); !ok {
		return "", code.CodeUserNotExist
	}
	// 2.check password
	if UserInformation.Password != utils.MD5(password) {
		return "", code.CodeInvalidPassword
	}
	// 3.return token
	token, err := myjwt.GenerateToken(UserInformation.ID, UserInformation.Username)
	if err != nil {
		return "", code.CodeServerBusy
	}
	return token, code.CodeSuccess
}

// Register 注册
func (u *UserDomainImpl) Register(ctx context.Context, email, password, captcha string) (string, code.Code) {
	var ok bool
	var UserInformation *entity.UserEntity

	// 1.判断用户是否存在
	if ok, _ = u.repo.IsExistUserByEmail(email); ok {
		return "", code.CodeUserExist
	}

	// 2. 从redis中判断验证码是否正确
	if ok, _ = redis.CheckCaptcha(email, captcha); !ok {
		return "", code.CodeInvalidCaptcha
	}

	// 3. gen username
	username := utils.GetRandomNumbers(11)

	// 4. store
	if UserInformation, ok = u.repo.Register(email, password, username); !ok {
		return "", code.CodeServerBusy
	}

	// 5.账户发送到邮箱账号，之后凭借账户登入
	if err := myemail.SendCaptcha(email, username, consts.UserNameMsg); err != nil {
		return "", code.CodeServerBusy
	}

	// 6. jwt token
	token, err := myjwt.GenerateToken(UserInformation.ID, UserInformation.Username)
	if err != nil {
		return "", code.CodeServerBusy
	}

	return token, code.CodeSuccess
}

// Captcha 发邮件
func (u *UserDomainImpl) Captcha(ctx context.Context, email string) code.Code {
	// 创建6为随机数
	send_code := utils.GetRandomNumbers(6)
	// 1. 存放到redis中
	if err := redis.SetCaptchaForEmail(email, send_code); err != nil {
		return code.CodeServerBusy
	}
	// 2.发送邮件
	if err := myemail.SendCaptcha(email, send_code, consts.CodeMsg); err != nil {
		return code.CodeServerBusy
	}
	return code.CodeSuccess
}
