package models

import (
	"github.com/lib/pq"
)

type Tag struct {
	ID      uint           `gorm:"primaryKey" json:"-"`
	EventID uint           `gorm:"index" json:"eventId"`
	Type    string         `gorm:"index:tags_type_idx" json:"type"`
	Values  pq.StringArray `gorm:"type:text[]" json:"values"`
}
