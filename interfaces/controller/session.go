package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai-ddd/application/session"
	logger "github.com/kangyueyue/go-ai-ddd/infrastructure/common/log"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/code"
	"github.com/kangyueyue/go-ai-ddd/interfaces/types/dto"
)

// GetUserSessionsByUserName 获取用户的会话列表
func GetUserSessionsByUserName(c *gin.Context) {
	res := new(dto.GetUserSessionsByUserNameResp)
	userName := c.GetString("userName") // from JWT middleware
	userSessions, err := session.GlobalSessionServiceImpl.GetUserSessionsByUserName(userName)

	if err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeServerBusy))
		return
	}

	res.Success()
	res.Sessions = userSessions
	c.JSON(http.StatusOK, res)
}

// ChatSend 聊天发送消息
func ChatSend(c *gin.Context) {
	res := new(dto.ChatSendResp)
	req := new(dto.ChatSendReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userName := c.GetString("userName") // from JWT middleware
	logger.Log.Infof("userName:%s", userName)
	aiInformation, code_ := session.GlobalSessionServiceImpl.ChatSend(userName, req.UserQuestion,
		req.ModelType, req.SessionID)
	if code_ != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}
	res.Success() // success
	res.AiInformation = aiInformation
	c.JSON(http.StatusOK, res)
}

// CreateSessionAndSendMessage 创建会话并发送消息
func CreateSessionAndSendMessage(c *gin.Context) {
	req := new(dto.CreateSessionAndSendMessageReq)
	res := new(dto.CreateSessionAndSendMessageResp)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userName := c.GetString("userName") // from JWT middleware
	logger.Log.Infof("userName:%s", userName)
	session_id, aiInformation, code_ := session.GlobalSessionServiceImpl.CreateSessionAndSendMessage(userName, req.UserQuestion, req.ModelType)
	if code_ != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}
	res.Success()
	res.AiInformation = aiInformation
	res.SessionID = session_id
	c.JSON(http.StatusOK, res)
}

// CreateStreamSessionAndSendMessage 创建流式会话并发送消息
func CreateStreamSessionAndSendMessage(c *gin.Context) {
	req := new(dto.CreateSessionAndSendMessageReq)
	userName := c.GetString("userName") // From JWT middleware
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid parameters"})
		return
	}
	// 设置SSE头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Accel-Buffering", "no") // 禁止代理缓存

	// 先创建会话并立即把 sesionId 下发给前端，随后开始流式输出
	sessionID, code_ := session.GlobalSessionServiceImpl.CreateStreamSessionOnly(userName, req.UserQuestion)
	if code_ != code.CodeSuccess {
		c.SSEvent("error", gin.H{"message": "Failed to create session"})
		return
	}
	// 先把sessionID 通过data事件发送给前端，前端据此绑定当前回答，侧边栏即可出现新标签
	c.Writer.WriteString(fmt.Sprintf("data: {\"sessionId\": \"%s\"}\n\n", sessionID))
	c.Writer.Flush()

	// 然后把本次回答进行流式发送
	code_ = session.GlobalSessionServiceImpl.SendMessageToExistSession(userName, sessionID, req.UserQuestion, req.ModelType, http.ResponseWriter(c.Writer))
	if code_ != code.CodeSuccess {
		c.SSEvent("error", gin.H{"message": "Failed to send message"})
		return
	}
}

// ChatHistory 聊天历史记录
func ChatHistory(c *gin.Context) {
	req := new(dto.ChatHistoryReq)
	res := new(dto.ChatHistoryResp)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, res.CodeOf(code.CodeInvalidParams))
		return
	}
	userName := c.GetString("userName") // from JWT middleware
	history, code_ := session.GlobalSessionServiceImpl.ChatHistory(userName, req.SessionID)
	if code_ != code.CodeSuccess {
		c.JSON(http.StatusOK, res.CodeOf(code_))
		return
	}
	res.Success()
	res.History = history
	c.JSON(http.StatusOK, res)
}

// ChatStreamSend 流式聊天发送消息
func ChatStreamSend(c *gin.Context) {
	req := new(dto.ChatSendReq)
	userName := c.GetString("userName") // From JWT middleware
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid parameters"})
		return
	}
	// 设置SSE头
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("X-Accel-Buffering", "no") // 禁止代理缓存

	// 然后把本次回答进行流式发送
	code_ := session.GlobalSessionServiceImpl.ChatSteamSend(userName, req.SessionID, req.UserQuestion, req.ModelType, http.ResponseWriter(c.Writer))
	if code_ != code.CodeSuccess {
		c.SSEvent("error", gin.H{"message": "Failed to send message"})
		return
	}
}
