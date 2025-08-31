package models

import (
	"time"
)

type PublicChatMessage struct {
	ID        uint               `gorm:"primaryKey" json:"id"`
	AccountID uint               `gorm:"index" json:"accountId"`
	Account   *Account           `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account"`
	ChannelID uint               `gorm:"index" json:"channelId"`
	Channel   *PublicChatChannel `gorm:"foreignKey:ChannelID;constraint:OnDelete:CASCADE" json:"channel"`
	CreatedAt time.Time          `json:"createdAt"`
	Content   string             `gorm:"type:text" json:"content"`
	EventID   uint               `gorm:"index" json:"eventId"`
	Event     *Event             `gorm:"foreignKey:EventID" json:"event"`
	NostrID   string             `gorm:"index" json:"nostrId"`
	PubKey    string             `gorm:"type:text;index" json:"pubkey"`
	Sig       string             `gorm:"type:text" json:"sig"`
	Tags      string             `gorm:"type:jsonb" json:"tags"`
}
