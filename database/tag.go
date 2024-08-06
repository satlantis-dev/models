package database

import (
	"github.com/lib/pq"
)

type Tag struct {
	ID      uint           `gorm:"primaryKey" json:"-"`
	EventID uint           `gorm:"index" json:"eventId"`
	Type    string         `json:"type"`
	Values  pq.StringArray `gorm:"type:text[]" json:"values"`
}
