package models

import (
	"time"

	"gorm.io/gorm"
)

type Calendar struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Slug        string          `gorm:"size:70" json:"slug"`
	Banner      string          `json:"banner"`
	AccountID   uint            `json:"account_id"`
	Account     *Account        `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	Events      []CalendarEvent `gorm:"many2many:calendar_calendar_events;" json:"events"`
	IsPublic    bool            `gorm:"default:true" json:"isPublic"`
	CreatedAt   time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"-"`
}
