package user

import (
	"errors"

	"github.com/kangyueyue/go-ai-ddd/domain/user/entity"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/utils"
	"gorm.io/gorm"
)

// UserRepositoryImpl 实现 Repository 接口
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository 创建 Repository 实例
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

// IsExistUserByEmail 判断用户是否存在
func (r *UserRepositoryImpl) IsExistUserByEmail(email string) (bool, *entity.UserEntity) {
	var user UserPojo
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在
		return false, nil
	}
	entity := Pojo2Entity(&user)
	return true, entity // 存在
}

// IsExistUserByUsername 判断用户是否存在
func (r *UserRepositoryImpl) IsExistUserByUsername(username string) (bool, *entity.UserEntity) {
	var user UserPojo
	err := r.db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在
		return false, nil
	}
	entity := Pojo2Entity(&user)
	return true, entity // 存在
}

// Register  注册
func (r *UserRepositoryImpl) Register(email, password, username string) (*entity.UserEntity, bool) {
	user := &UserPojo{
		Email:    email,
		Name:     username,
		Password: utils.MD5(password), // 加密存储
		Username: username,
	}
	err := r.db.Create(user).Error
	if err != nil {
		return nil, false
	}
	entity := Pojo2Entity(user)
	return entity, true
}
