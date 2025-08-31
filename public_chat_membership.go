package models

import "time"

type PublicChatMembership struct {
	ID                  uint               `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time          `json:"-"`
	UpdatedAt           time.Time          `json:"-"`
	DeletedAt           *time.Time         `gorm:"index" json:"-,omitempty"`
	AccountID           uint               `gorm:"index" json:"accountId"`
	Account             *Account           `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account"`
	LastReadNoteID      *uint              `json:"lastReadNoteId"`
	PublicChatChannelID uint               `gorm:"index" json:"publicChatChannelId"`
	PublicChatChannel   *PublicChatChannel `gorm:"foreignKey:PublicChatChannelID;constraint:OnDelete:CASCADE" json:"publicChatChannel"`
}
