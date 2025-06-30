package models

import "time"

type Block struct {
	BlockerID     uint      `gorm:"index:idx_blocker_blocking,unique"`
	BlockingID    uint      `gorm:"index:idx_blocker_blocking,unique"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	Bidirectional bool      `gorm:"default:false" json:"bidirectional"` // Bidirectional -> block, Unidirectional -> mute
}
