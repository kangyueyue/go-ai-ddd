package message

import (
	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
	"gorm.io/gorm"
)

// MessageRepository
type MessageRepository struct {
	db *gorm.DB
}

// NewMessageRepository 创建消息仓库
func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

// GetAllMessages 获取所有消息
func (m *MessageRepository) GetAllMessages() ([]MessagePojo, error) {
	var msgs []MessagePojo
	err := m.db.Order("created_at asc").Find(&msgs).Error
	return msgs, err
}

// CreateMessage 创建消息
func (m *MessageRepository) CreateMessage(message *entity.MessageEntity) (*entity.MessageEntity, error) {
	pojo := EntityToPojo(message)
	err := m.db.Create(&pojo).Error
	return PojoToEntity(pojo), err
}
