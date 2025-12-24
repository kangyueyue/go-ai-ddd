package message

import "github.com/kangyueyue/go-ai-ddd/domain/session/entity"


func EntityToPojo(entity *entity.MessageEntity) *MessagePojo {
	pojo := &MessagePojo{
		ID:        entity.ID,
		SessionID: entity.SessionID,
		UserName:  entity.UserName,
		Content:   entity.Content,
		IsUser:    entity.IsUser,
	}
	return pojo
}


func PojoToEntity(pojo *MessagePojo) *entity.MessageEntity {
	entity := &entity.MessageEntity{
		ID:        pojo.ID,
		SessionID: pojo.SessionID,
		UserName:  pojo.UserName,
		Content:   pojo.Content,
		IsUser:    pojo.IsUser,
	}
	return entity
}