package session

import (
	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
	"gorm.io/gorm"
)

// SessionRepositoryImpl 实现 Repository 接口
type SessionRepositoryImpl struct {
	db *gorm.DB
}


// NewSessionRepositoryImpl 创建用户仓库
func NewSessionRepositoryImpl(db *gorm.DB) *SessionRepositoryImpl {
	return &SessionRepositoryImpl{db: db}
}

// CreateSession 创建会话
func (r *SessionRepositoryImpl) CreateSession(session *entity.SessionEntity) (*entity.SessionEntity, error) {
	pojo := Entity2Pojo(session)
	err := r.db.Create(pojo).Error
	return Pojo2Entity(pojo), err
}

// GetSessionInfosBySessionIDs 获取会话信息
func (r *SessionRepositoryImpl) GetSessionInfosBySessionIDs(sessionIDs []string) ([]entity.SessionInfoEntity, error) {
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
