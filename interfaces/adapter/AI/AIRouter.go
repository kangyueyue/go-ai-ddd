package AI

import (
	"github.com/gin-gonic/gin"
	"github.com/kangyueyue/go-ai-ddd/interfaces/controller"
)

// AIRouter sets up AI related routes
func AIRouter(r *gin.RouterGroup) {
	r.GET("/chat/sessions", controller.GetUserSessionsByUserName)
	r.POST("/chat/send-new-session", controller.CreateSessionAndSendMessage)
	r.POST("/chat/send", controller.ChatSend)
	r.POST("/chat/history", controller.ChatHistory)
	// r.POST("/chat/tts", AI.ChatSpeech)                  // ChatSpeechHandler
	r.POST("/chat/send-stream-new-session", controller.CreateStreamSessionAndSendMessage)
	r.POST("/chat/send-stream", controller.ChatStreamSend)
}
