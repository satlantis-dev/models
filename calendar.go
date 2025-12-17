package models

import (
	"time"

	"gorm.io/gorm"
)

type Calendar struct {
	ID               uint               `json:"id"`
	AccountID        uint               `json:"account_id"`
	Account          *Account           `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	Slug             string             `gorm:"size:70" json:"slug"`
	Banner           string             `json:"banner"`
	CreatedAt        time.Time          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt        *gorm.DeletedAt    `gorm:"index" json:"-"`
	Events           []CalendarEvent    `gorm:"many2many:calendar_calendar_events;constraint:OnDelete:CASCADE;" json:"events,omitempty"`
	EventCount       int                `gorm:"-" json:"eventCount"`
	EventTags        []CalendarEventTag `gorm:"many2many:calendar_calendar_event_tags;constraint:OnDelete:CASCADE" json:"eventTags,omitempty"`
	Contributors     *[]AccountMiniDTO  `gorm:"-" json:"contributors,omitempty"`
	NumSubscriptions int                `gorm:"-" json:"numSubscriptions"`
	PlaceID          *uint              `json:"placeId,omitempty"`
	Place            *Place             `gorm:"foreignKey:PlaceID;constraint:OnDelete:SET NULL" json:"place,omitempty"`
	IsPublic         bool               `gorm:"default:true" json:"isPublic"`
	Featured         bool               `gorm:"default:false" json:"featured"`
}
