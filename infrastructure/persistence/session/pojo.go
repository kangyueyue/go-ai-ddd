package session

import (
	"time"

	"gorm.io/gorm"
)

// Session 会话
type SessionPojo struct {
	ID        string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserName  string         `gorm:"index;not null" json:"username"`
	Title     string         `gorm:"type:varchar(100)" json:"title"`
	CreatedAt time.Time      `json:"created_at" ` // 数据库自动填充当前时间
	UpdatedAt time.Time      `json:"updated_at"`  // 数据库自动更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// SessionInfo 会话信息
type SessionInfo struct {
	SessionID string `json:"sessionId"`
	Title     string `json:"name"`
}

func (s *SessionPojo) TableName() string {
	return "sessions"
}
