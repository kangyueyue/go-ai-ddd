package repository

import "github.com/kangyueyue/go-ai-ddd/domain/session/entity"

type ISessionRepository interface {
	CreateSession(session *entity.SessionEntity) (*entity.SessionEntity, error)
	GetSessionInfosBySessionIDs(sessionIDs []string) ([]entity.SessionInfoEntity, error)
}
