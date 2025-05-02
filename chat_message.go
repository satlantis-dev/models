package models

import "time"

type ChatMessage struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	NostrID        string       `gorm:"index" json:"nostrId"`
	ConversationID uint         `json:"conversationId"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID" json:"conversation"`
	PushSentAt     *time.Time   `json:"pushSentAt"`
	CreatedAt      time.Time    `json:"-"`
	DeletedAt      *time.Time   `json:"-"`
}
