package models

import (
	"time"
)

type PublicChatMessage struct {
	ID        uint              `gorm:"primaryKey" json:"id"`
	AccountID uint              `gorm:"index" json:"accountId"`
	Account   AccountDTO        `json:"account" gorm:"foreignKey:AccountID"`
	ChannelID uint              `gorm:"index" json:"channelId"`
	Channel   PublicChatChannel `json:"channel" gorm:"foreignKey:ChannelID"`
	CreatedAt time.Time         `json:"createdAt"`
	Content   string            `gorm:"type:text" json:"content"`
	EventID   uint              `gorm:"index" json:"eventId"`
	Event     Event             `json:"event"`
	NostrID   string            `gorm:"index" json:"nostrId"`
	PubKey    string            `gorm:"type:text;index" json:"pubkey"`
	Sig       string            `gorm:"type:text" json:"sig"`
	Tags      string            `gorm:"type:jsonb" json:"tags"`
}
