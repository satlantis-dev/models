package models

import "github.com/nbd-wtf/go-nostr"

type Event struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	NostrID   string `gorm:"uniqueIndex" json:"nostrId"`
	CreatedAt int64  `json:"createdAt"`
	Content   string `gorm:"type:text" json:"content"`
	Kind      uint   `gorm:"index" json:"kind"`
	PubKey    string `gorm:"type:text;index" json:"pubkey"`
	Sig       string `gorm:"type:text" json:"sig"`
	//Tags       []Tag      `gorm:"foreignKey:EventID" json:"tags"`
	Tags       []Tag      `gorm:"foreignKey:EventID;constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;" json:"tags"`
	TagsRaw    nostr.Tags `gorm:"type:jsonb;serializer:json" json:"tagsData"`
	Reconciled bool       `gorm:"default:false" json:"reconciled"`
}
