package models

import "time"

type PublicChatMembership struct {
	ID                  uint              `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time         `json:"-"`
	UpdatedAt           time.Time         `json:"-"`
	DeletedAt           *time.Time        `gorm:"index" json:"-,omitempty"`
	AccountID           uint              `gorm:"index" json:"accountId"`
	Account             Account           `json:"account"`
	LastReadNoteID      *uint             `json:"lastReadNoteId"`
	PublicChatChannelID uint              `gorm:"index" json:"publicChatChannelId"`
	PublicChatChannel   PublicChatChannel `json:"publicChatChannel"`
}
