package models

import (
	"time"

	"gorm.io/gorm"
)

type CalendarEventTag struct {
	ID        uint            `gorm:"primaryKey" json:"id,omitempty"`
	Name      string          `gorm:"type:text;not null;uniqueIndex" json:"name,omitempty"`
	NumEvents *int            `gorm:"default:0" json:"numEvents,omitempty"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}

func (CalendarEventTag) TableName() string {
	return "calendar_event_tags"
}
