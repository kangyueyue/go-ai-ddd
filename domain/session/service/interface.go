package session

import (
	"net/http"

	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
)

// IServiceDomain 服务领域接口
type IServiceDomain interface {
	GetUserSessionsByUserName(userName string) ([]entity.SessionInfoEntity, error)
	ChatSend(userName string, userQuestion string, modelType string, sessionID string) (string, code.Code)
	ChatHistory(userName, sessionId string) ([]entity.HistoryEntity, code.Code)
	CreateSessionAndSendMessage(userName, userQuestion, modelType string) (string, string, code.Code)
	CreateStreamSessionOnly(userName string, userQuestion string) (string, code.Code)
	SendMessageToExistSession(userName, sessionID, userQuestion, modelType string, writer http.ResponseWriter) code.Code
	ChatSteamSend(userName, sessionID, userQuestion, modelType string, writer http.ResponseWriter) code.Code
}
