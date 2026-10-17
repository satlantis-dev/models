package models

import "github.com/nbd-wtf/go-nostr"

type Event struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	NostrID    string     `gorm:"index" json:"nostrId"`
	CreatedAt  int64      `gorm:"index:idx_events_unreconciled_kind_created_at,priority:2,where:reconciled = false,option:CONCURRENTLY;index:idx_events_pub_key_kind_created_at,priority:3,option:CONCURRENTLY" json:"createdAt"`
	Content    string     `gorm:"type:text" json:"content"`
	Kind       uint       `gorm:"index;index:idx_events_unreconciled_kind_created_at,priority:1,where:reconciled = false,option:CONCURRENTLY;index:idx_events_pub_key_kind_created_at,priority:2,option:CONCURRENTLY" json:"kind"`
	PubKey     string     `gorm:"type:text;index;index:idx_events_pub_key_kind_created_at,priority:1,option:CONCURRENTLY" json:"pubkey"`
	Sig        string     `gorm:"type:text" json:"sig"`
	Tags       []Tag      `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"tags"`
	TagsRaw    nostr.Tags `gorm:"type:jsonb;serializer:json" json:"tagsData"`
	Reconciled bool       `gorm:"default:false" json:"reconciled"`
}
