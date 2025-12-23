package entity

// MessageEntity 消息
type MessageEntity struct {
	ID        uint   `json:"id"`
	SessionID string `json:"session_id"`
	UserName  string `json:"username"`
	Content   string `json:"content"`
	IsUser    bool   `json:"is_user"`
}

// HistoryEntity 历史消息
type HistoryEntity struct {
	IsUser  bool   `json:"is_user"`
	Content string `json:"content"`
}
