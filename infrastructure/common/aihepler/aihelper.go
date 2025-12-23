package aihelper

import (
	"context"
	"sync"

	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/mq"
	"github.com/kangyueyue/go-ai-ddd/infrastructure/utils"
)

// AIHelper ai helper
type AIHelper struct {
	model    AIModel
	messages []*entity.MessageEntity
	mu       sync.RWMutex
	// 一个会话绑定一个
	SessionID string
	saveFunc  func(*entity.MessageEntity) (*entity.MessageEntity, error)
}

// NewAIHelper 创建一个ai helper
func NewAIHelper(model_ AIModel, sessionID string) *AIHelper {
	return &AIHelper{
		model:     model_,
		messages:  make([]*entity.MessageEntity, 0),
		mu:        sync.RWMutex{},
		SessionID: sessionID,
		saveFunc: func(msg *entity.MessageEntity) (*entity.MessageEntity, error) {
			// 保存策略是通过m异步q保存到db中
			data := mq.GenerateMessageMQParam(msg.SessionID, msg.Content, msg.UserName, msg.IsUser)
			err := mq.RMQMessage.Publish(string(data)) // publish 生产者
			return msg, err
		},
	}
}

// SetSaveFunc 设置保存函数
func (a *AIHelper) SetSaveFunc(saveFunc func(*entity.MessageEntity) (*entity.MessageEntity, error)) {
	a.saveFunc = saveFunc
}

// addMessage 添加消息
func (a *AIHelper) AddMessage(content, userName string,
	isUser bool, save bool,
) {
	userMsg := entity.MessageEntity{
		SessionID: a.SessionID,
		Content:   content,
		UserName:  userName,
		IsUser:    isUser,
	}
	a.messages = append(a.messages, &userMsg)
	if save {
		// 是否需要持久化
		a.saveFunc(&userMsg)
	}
}

// GenrateResponse 生成响应
func (a *AIHelper) GenerateResponse(userName string, ctx context.Context,
	userQuestion string,
) (*entity.MessageEntity, error) {
	// 调用存储函数
	a.AddMessage(userQuestion, userName, true, true)
	a.mu.Lock()

	// 将entity.MessageEntity转化为Enio接受的数据类型
	messages := utils.ConvertToSchemaMessages(a.messages)
	a.mu.Unlock()

	// 调用模型生成回答结果
	schemaMsg, err := a.model.GenerateResponse(ctx, messages)
	if err != nil {
		return nil, err
	}
	// 将schema.Message转化为entity.MessageEntity
	modelMsg := utils.ConvertToModelMessages(a.SessionID, userName, schemaMsg)

	// save
	a.saveFunc(modelMsg)

	return modelMsg, nil
}

// GetHistory 获取历史消息
func (a *AIHelper) GetHistory() []*entity.MessageEntity {
	a.mu.RLock()
	defer a.mu.RUnlock()
	res := make([]*entity.MessageEntity, len(a.messages))
	copy(res, a.messages)
	return res
}

// StreamResponse 流式响应
func (a *AIHelper) StreamResponse(userName string, ctx context.Context,
	cb StreamCallback, userQuestion string) (*entity.MessageEntity, error) {
	// 调用存储函数，保存用户问题
	a.AddMessage(userQuestion, userName, true, true)
	a.mu.RLock()
	messages := utils.ConvertToSchemaMessages(a.messages)
	a.mu.RUnlock()

	content, err := a.model.StreamResponse(ctx, messages, cb)
	if err != nil {
		return nil, err
	}
	modelMsg := &entity.MessageEntity{
		SessionID: a.SessionID,
		UserName:  userName,
		Content:   content,
		IsUser:    false,
	}
	// 调用存储函数，保存AI回答
	a.saveFunc(modelMsg)
	return modelMsg, nil
}
