package user

import "github.com/kangyueyue/go-ai-ddd/domain/user/entity"

// Entity2Pojo entity 转 pojo
func Entity2Pojo(user *entity.UserEntity) *UserPojo {
	return &UserPojo{
		Username: user.Username,
		Password: user.Password,
	}
}

// Pojo2Entity pojo 转 entity
func Pojo2Entity(user *UserPojo) *entity.UserEntity {
	return &entity.UserEntity{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}
}
