package models

import "time"

type Save struct {
	AccountID    uint      `gorm:"index:idx_save,unique"`
	CollectionID uint      `gorm:"index:idx_save,unique"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}
