package entity

// SessionEntity 会话
type SessionEntity struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Title    string `json:"title"`
}

// SessionInfoEntity 会话信息
type SessionInfoEntity struct {
	SessionID string `json:"sessionId"`
	Title     string `json:"name"`
}
