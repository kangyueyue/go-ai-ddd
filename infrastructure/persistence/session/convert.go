package session

import "github.com/kangyueyue/go-ai-ddd/domain/session/entity"

func Entity2Pojo(entity *entity.SessionEntity) *SessionPojo {
	return &SessionPojo{
		ID:       entity.ID,
		UserName: entity.UserName,
		Title:    entity.Title,
	}
}

func Pojo2Entity(pojo *SessionPojo) *entity.SessionEntity {
	return &entity.SessionEntity{
		ID:       pojo.ID,
		UserName: pojo.UserName,
		Title:    pojo.Title,
	}
}
