package models

import "time"

type Block struct {
	BlockerID     uint      `gorm:"primaryKey;index:idx_blocker_blocking,unique;uniqueIndex:idx_blockerid"`
	BlockingID    uint      `gorm:"primaryKey;index:idx_blocker_blocking,unique;uniqueIndex:idx_blockingid"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	Bidirectional bool      `gorm:"default:false" json:"bidirectional"` // Bidirectional -> block, Unidirectional -> mute
}
