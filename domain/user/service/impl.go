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
	"github.com/kangyueyue/go-ai-ddd/interfaces/types"
)

// UserDomainImpl 用户领域实现
type UserDomainImpl struct {
	repo repository.IUserRepository
}

// NewUserDomain
func NewUserDomainImpl(repo repository.IUserRepository) IUserDomain {
	return &UserDomainImpl{repo: repo}
}

func (u *UserDomainImpl) Login(ctx context.Context, email, passwd, captcha string) {

}

// Register 注册
func (u *UserDomainImpl) Register(ctx context.Context, email, password, captcha string) (string, types.Code) {
	var ok bool
	var UserInformation *entity.UserEntity

	// 1.判断用户是否存在
	if ok, _ = u.repo.IsExistUserByEmail(email); ok {
		return "", types.CodeUserExist
	}

	// 2. 从redis中判断验证码是否正确
	if ok, _ = redis.CheckCaptcha(email, captcha); !ok {
		return "", types.CodeInvalidCaptcha
	}

	// 3. gen username
	username := utils.GetRandomNumbers(11)

	// 4. store
	if UserInformation, ok = u.repo.Register(email, password, username); !ok {
		return "", types.CodeServerBusy
	}

	// 5.账户发送到邮箱账号，之后凭借账户登入
	if err := myemail.SendCaptcha(email, username, consts.UserNameMsg); err != nil {
		return "", types.CodeServerBusy
	}

	// 6. jwt token
	token, err := myjwt.GenerateToken(UserInformation.ID, UserInformation.Username)
	if err != nil {
		return "", types.CodeServerBusy
	}

	return token, types.CodeSuccess
}

func (u *UserDomainImpl) Captcha(ctx context.Context) {

}
