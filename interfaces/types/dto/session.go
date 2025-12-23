package dto

import (
	"github.com/kangyueyue/go-ai-ddd/domain/session/entity"
)

type (
	GetUserSessionsByUserNameReq struct {
		UserQuestion string `json:"question" binding:"required"`
		ModelType    string `json:"modelType" binding:"required"`
	}
	GetUserSessionsByUserNameResp struct {
		Response
		Sessions []entity.SessionInfoEntity `json:"sessions,omitempty"`
	}
	ChatSendReq struct {
		UserQuestion string `json:"question" binding:"required"`
		ModelType    string `json:"modelType" binding:"required"`
		SessionID    string `json:"sessionId,omitempty" binding:"required"`
	}
	ChatSendResp struct {
		Response
		// AI回答
		AiInformation string `json:"Information,omitempty"`
	}
	ChatHistoryReq struct {
		SessionID string `json:"sessionId,omitempty" binding:"required"`
	}
	ChatHistoryResp struct {
		Response
		History []entity.HistoryEntity `json:"history"`
	}
	CreateSessionAndSendMessageReq struct {
		UserQuestion string `json:"question" binding:"required"`
		ModelType    string `json:"modelType" binding:"required"`
	}
	CreateSessionAndSendMessageResp struct {
		Response
		// AI回答
		AiInformation string `json:"Information,omitempty"`
		// 当前会话ID
		SessionID string `json:"sessionId,omitempty"`
	}
)
