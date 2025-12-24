package session

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
	"github.com/kangyueyue/go-ai-ddd/domain/session/repository"
	aihelper "github.com/kangyueyue/go-ai-ddd/infrastructure/common/aihepler"
	logger "github.com/kangyueyue/go-ai-ddd/infrastructure/common/log"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
)

var ctx = context.Background()

// SessionDomainImpl 会话领域实现
type SessionDomainImpl struct {
	sessionRepository repository.ISessionRepository
}

// NewSessionDomainImpl 创建会话领域实现
func NewSessionDomainImpl(sessionRepository repository.ISessionRepository) IServiceDomain {
	return &SessionDomainImpl{
		sessionRepository: sessionRepository,
	}
}

// GetUserSessionsByUserName 获取用户的会话列表
func (s *SessionDomainImpl) GetUserSessionsByUserName(userName string) ([]entity.SessionInfoEntity, error) {
	manager := aihelper.GetGlobalManager() // 获得全局的aihelper管理器
	Sessions := manager.GetAllSessionID(userName)

	return s.sessionRepository.GetSessionInfosBySessionIDs(Sessions)
}

// ChatSend 聊天发送消息
func (s *SessionDomainImpl) ChatSend(userName string, userQuestion string,
	modelType string, sessionID string,
) (string, code.Code) {
	// 获取helper
	manager := aihelper.GetGlobalManager()
	config := map[string]interface{}{
		"apiKey": "your-api-key", // TODO: 从配置中获取
	}
	helper, err := manager.GetOrCreateAIHelper(userName, sessionID, modelType, config)
	if err != nil {
		return "", code.AIModelFail
	}
	// 生成ai回答
	res, err := helper.GenerateResponse(userName, ctx, userQuestion)
	if err != nil {
		logger.Log.Errorf("ChatSend GenerateResponse error: %v", err)
		return "", code.AIModelFail
	}
	return res.Content, code.CodeSuccess
}

// ChatHistory 聊天历史记录
func (s *SessionDomainImpl) ChatHistory(userName, sessionId string) ([]entity.HistoryEntity, code.Code) {
	// helper中的消息历史
	manager := aihelper.GetGlobalManager()
	helper, ok := manager.GetAIHelper(userName, sessionId)
	if !ok {
		return nil, code.CodeServerBusy
	}
	message := helper.GetHistory()
	history := make([]entity.HistoryEntity, 0, len(message))

	// 转化为为历史消息
	for i, msg := range message {
		isUser := i%2 == 0 // 是否为用户还是AI回答
		history = append(history, entity.HistoryEntity{
			IsUser:  isUser,
			Content: msg.Content,
		})
	}
	return history, code.CodeSuccess
}

// CreateSessionAndSendMessage 创建会话并发送消息
func (s *SessionDomainImpl) CreateSessionAndSendMessage(
	userName, userQuestion, modelType string,
) (string, string, code.Code) {
	// 创建一个新的会话
	newSession := &entity.SessionEntity{
		ID:       uuid.New().String(),
		UserName: userName,
		Title:    userQuestion,
	}
	createSession, err := s.sessionRepository.CreateSession(newSession)
	if err != nil {
		logger.Log.Errorf("CreateSessionAndSendMessage CreateSession error: %v", err)
		return "", "", code.CodeServerBusy
	}

	manager := aihelper.GetGlobalManager()
	config := map[string]interface{}{
		"apiKey": "your-api-key", // TODO: 从配置中获取
	}
	helper, err := manager.GetOrCreateAIHelper(userName, createSession.ID, modelType, config)
	if err != nil {
		return "", "", code.AIModelFail
	}
	// 生成AI回答
	res, err := helper.GenerateResponse(userName, ctx, userQuestion)
	if err != nil {
		logger.Log.Errorf("ChatSend GenerateResponse error: %v", err)
		return "", "", code.AIModelFail
	}
	return createSession.ID, res.Content, code.CodeSuccess
}

// CreateStreamSessionOnly 创建流式会话
func (s *SessionDomainImpl) CreateStreamSessionOnly(userName string, userQuestion string) (string, code.Code) {
	newSession := &entity.SessionEntity{
		ID:       uuid.New().String(),
		UserName: userName,
		Title:    userQuestion,
	}
	createdSession, err := s.sessionRepository.CreateSession(newSession)
	if err != nil {
		logger.Log.Errorf("CreateStreamSessionOnly CreateSession error: %v", err)
		return "", code.CodeServerBusy
	}
	return createdSession.ID, code.CodeSuccess
}

// SendMessageToExistSession 发送消息到已存在的会话
func (s *SessionDomainImpl) SendMessageToExistSession(userName, sessionID,
	userQuestion, modelType string, writer http.ResponseWriter,
) code.Code {
	// 确保writer 支持flush
	flusher, ok := writer.(http.Flusher)
	if !ok {
		logger.Log.Errorf("SendMessageToExistSession wirter not support flush")
		return code.CodeServerBusy
	}

	manager := aihelper.GetGlobalManager()
	config := map[string]interface{}{
		"apiKey": "your-api-key", // TODO: 从配置中获取
	}
	helper, err := manager.GetOrCreateAIHelper(userName, sessionID, modelType, config)
	if err != nil {
		logger.Log.Errorf("SendMessageToExistSession GetOrCreateAIHelper error: %v", err)
		return code.AIModelFail
	}
	// 回调函数
	cb := func(msg string) {
		// 直接发送数据，不转义
		logger.Log.Infof("[SSE] Sending chunk: %s(len: %d)", msg, len(msg))
		_, err := writer.Write([]byte("data: " + msg + "\n\n"))
		if err != nil {
			logger.Log.Errorf("[SSE] Write error: %v", err)
			return
		}
		flusher.Flush() // 每次必须flush
		logger.Log.Infof("[SSE] Flushed chunk: %s(len: %d)", msg, len(msg))
	}
	_, err_ := helper.StreamResponse(userName, ctx, cb, userQuestion)
	if err_ != nil {
		logger.Log.Errorf("SendMessageToExistSession StreamResponse error: %v", err)
		return code.AIModelFail
	}

	_, err = writer.Write([]byte("data: [DONE]\n\n"))
	if err != nil {
		logger.Log.Errorf("[SSE] Write error: %v", err)
		return code.CodeServerBusy
	}
	flusher.Flush()
	return code.CodeSuccess
}

// ChatSteamSend 聊天流式发送消息
func (s *SessionDomainImpl) ChatSteamSend(userName, sessionID, userQuestion, modelType string, writer http.ResponseWriter) code.Code {
	return s.SendMessageToExistSession(userName, sessionID, userQuestion, modelType, writer)
}
