package models

import (
	"time"
)

type MetricSource struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Frequency   string    `gorm:"type:text" json:"frequency"`
	LastUpdated time.Time `json:"-"`
	Name        string    `gorm:"type:text" json:"name"`
	URL         string    `gorm:"type:text" json:"url"`
}
