package persistence

import (
	sessionDB "github.com/kangyueyue/go-ai-ddd/domain/session/repository"
	"github.com/kangyueyue/go-ai-ddd/domain/user/repository"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/session"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/persistence/user"
	"gorm.io/gorm"
)

// 所有的仓储
type Repositories struct {
	User repository.IUserRepository
	Session sessionDB.ISessionRepository
}

// NewRepositories 初始化仓储层
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		User: user.NewUserRepository(db),
		Session: session.NewSessionRepositoryImpl(db),
	}
}