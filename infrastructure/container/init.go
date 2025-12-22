package container

import (
	"github.com/kangyueyue/go-ai-ddd/application/user"
	userSrv "github.com/kangyueyue/go-ai-ddd/domain/user/service"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence"
	mysql "github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/db"
)

// LoadingDomain 加载领域层
func LoadingDomain() {
	// repos
	repos := persistence.NewRepositories(mysql.DB)

	// user domain
	// 依赖关系 user domain 依赖 user repository
	userDomain := userSrv.NewUserDomainImpl(repos.User)
	user.GetUserServiceImpl(userDomain)
}
