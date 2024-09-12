package models

import "time"

type PlacePublicChatChannel struct {
	ID                  uint              `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time         `json:"-"`
	UpdatedAt           time.Time         `json:"-"`
	DeletedAt           *time.Time        `gorm:"index" json:"-,omitempty"`
	PlaceID             uint              `gorm:"index" json:"placeId"`
	PublicChatChannelID uint              `gorm:"index" json:"publicChatChannelId"`
	PublicChatChannel   PublicChatChannel `json:"publicChatChannel"`
}
