package session

import (
	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
	"gorm.io/gorm"
)

// UserRepositoryImpl 实现 Repository 接口
type UserRepositoryImpl struct {
	db *gorm.DB
}

// CreateSession 创建会话
func (r *UserRepositoryImpl) CreateSession(session *entity.SessionEntity) (*entity.SessionEntity, error) {
	pojo := Entity2Pojo(session)
	err := r.db.Create(pojo).Error
	return Pojo2Entity(pojo), err
}

// GetSessionInfosBySessionIDs 获取会话信息
func (r *UserRepositoryImpl) GetSessionInfosBySessionIDs(sessionIDs []string) ([]entity.SessionInfoEntity, error) {
	sessions := make([]SessionPojo, 0, len(sessionIDs))
	err := r.db.Model(&SessionPojo{}).Where("id IN (?)", sessionIDs).Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	sessionInfos := make([]entity.SessionInfoEntity, 0, len(sessions))
	for _, session := range sessions {
		sessionInfos = append(sessionInfos, entity.SessionInfoEntity{
			SessionID: session.ID,
			Title:     session.Title,
		})
	}
	return sessionInfos, nil
}
