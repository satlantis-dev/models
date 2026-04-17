package models

import (
	"time"

	"gorm.io/gorm"
)

type Theme struct {
	ID              uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string          `gorm:"not null;uniqueIndex" json:"name"`
	CreatedAt       time.Time       `gorm:"autoCreateTime" json:"-"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime" json:"-"`
	DeletedAt       *gorm.DeletedAt `gorm:"index" json:"-"`
	BackgroundColor string          `gorm:"type:varchar(7);not null" json:"backgroundColor"`
	ForegroundColor string          `gorm:"type:varchar(7);not null" json:"foregroundColor"`
	Rank            *int            `json:"rank,omitempty"`
}
