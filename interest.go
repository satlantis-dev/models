package models

import (
	"time"

	"github.com/lib/pq"
)

type Interest struct {
	ID                    uint           `gorm:"primaryKey" json:"id"`
	Name                  string         `gorm:"unique" json:"name"`
	CreatedAt             time.Time      `json:"-"`
	UpdatedAt             time.Time      `json:"-"`
	DeletedAt             *time.Time     `gorm:"index" json:"-,omitempty"`
	Description           string         `gorm:"type:text" json:"description"`
	RecommendationsByNpub pq.StringArray `gorm:"type:varchar[]" json:"recommendations_by_npub"`
	RecommendationsById   []uint         `gorm:"type:integer[]" json:"recommendations_by_id"`
	Hashtags              pq.StringArray `gorm:"type:varchar[]" json:"hashtags"`
}
