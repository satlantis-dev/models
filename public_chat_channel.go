package models

import (
	"time"

	"github.com/lib/pq"
)

type PublicChatChannel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	About     string         `json:"about"`
	AccountID uint           `gorm:"index" json:"accountId"`
	Account   *Account       `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account"`
	CreatedAt *time.Time     `json:"createdAt"`
	Content   string         `gorm:"type:text" json:"content"`
	EventID   uint           `gorm:"index" json:"eventId"`
	Event     *Event         `gorm:"foreignKey:EventID" json:"event"`
	Kind      uint           `gorm:"index" json:"kind"`
	Name      string         `json:"name"`
	NostrID   string         `gorm:"index" json:"nostrId"`
	PubKey    string         `gorm:"type:text;index" json:"pubkey"`
	Sig       string         `gorm:"type:text" json:"sig"`
	Tags      string         `gorm:"type:jsonb" json:"tags"`
	Picture   string         `json:"picture"`
	Relays    pq.StringArray `gorm:"type:text[]" json:"relays"`
}
