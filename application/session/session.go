package session

import (
	"net/http"
	"sync"

	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
	session "github.com/kangyueyue/go-ai-ddd/domain/session/service"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
)

// ISessionService 会话服务接口
type ISessionService interface {
	GetUserSessionsByUserName(userName string) ([]entity.SessionInfoEntity, error)
	ChatSend(userName string, userQuestion string, modelType string, sessionID string) (string, code.Code)
	ChatHistory(userName, sessionId string) ([]entity.HistoryEntity, code.Code)
	CreateSessionAndSendMessage(userName, userQuestion, modelType string) (string, string, code.Code)
	CreateStreamSessionOnly(userName string, userQuestion string) (string, code.Code)
	SendMessageToExistSession(userName, sessionID, userQuestion, modelType string, writer http.ResponseWriter) code.Code
	ChatSteamSend(userName, sessionID, userQuestion, modelType string, writer http.ResponseWriter) code.Code
}

// SessionService 会话服务实现
type SessionService struct {
	domain session.IServiceDomain
}

var (
	GlobalSessionServiceImpl *SessionService
	SessionServiceImplOnce   sync.Once
)

// GetGlobalSessionServiceImpl 获取全局的会话服务实例
func GetGlobalSessionServiceImpl(u session.IServiceDomain) ISessionService {
	if GlobalSessionServiceImpl != nil {
		return GlobalSessionServiceImpl
	}
	SessionServiceImplOnce.Do(func() {
		if GlobalSessionServiceImpl == nil {
			GlobalSessionServiceImpl = &SessionService{
				domain: u,
			}
		}
	})
	return GlobalSessionServiceImpl
}

func (s *SessionService) GetUserSessionsByUserName(userName string) ([]entity.SessionInfoEntity, error) {
	return s.domain.GetUserSessionsByUserName(userName)
}
func (s *SessionService) ChatSend(userName string, userQuestion string, modelType string, sessionID string) (string, code.Code) {
	return s.domain.ChatSend(userName, userQuestion, modelType, sessionID)
}
func (s *SessionService) ChatHistory(userName, sessionId string) ([]entity.HistoryEntity, code.Code) {
	return s.domain.ChatHistory(userName, sessionId)
}
func (s *SessionService) CreateSessionAndSendMessage(userName, userQuestion, modelType string) (string, string, code.Code) {
	return s.domain.CreateSessionAndSendMessage(userName, userQuestion, modelType)
}

func (s *SessionService) CreateStreamSessionOnly(userName string, userQuestion string) (string, code.Code) {
	return s.domain.CreateStreamSessionOnly(userName, userQuestion)
}
func (s *SessionService) SendMessageToExistSession(userName, sessionID, userQuestion, modelType string, writer http.ResponseWriter) code.Code {
	return s.domain.SendMessageToExistSession(userName, sessionID, userQuestion, modelType, writer)
}
func (s *SessionService) ChatSteamSend(userName, sessionID, userQuestion, modelType string, writer http.ResponseWriter) code.Code {
	return s.domain.ChatSteamSend(userName, sessionID, userQuestion, modelType, writer)
}
