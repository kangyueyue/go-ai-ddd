package repository

import "github.com/kangyueyue/go-ai-ddd/domain/user/entity"

// IUserRepository 用户仓储接口
type IUserRepository interface {
	IsExistUserByEmail(email string) (bool, *entity.UserEntity)
	IsExistUserByUsername(username string) (bool, *entity.UserEntity)
	Register(email, password, username string) (*entity.UserEntity, bool)
}
