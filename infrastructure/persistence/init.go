package persistence

import (
	"github.com/kangyueyue/go-ai-ddd/domain/user/repository"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/user"
	"gorm.io/gorm"
)

// 所有的仓储
type Repositories struct {
	User repository.IUserRepository
}

// NewRepositories 初始化仓储层
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewUserRepository(db),
	}
}
