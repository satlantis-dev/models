package models

import (
	"time"

	"gorm.io/gorm"
)

type CalendarEventNote struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	CalendarEventID uint            `json:"calendarEventId"`
	NoteID          uint            `json:"noteId"`
	Note            *Note           `gorm:"constraint:OnDelete:CASCADE;" json:"note,omitempty"`
	CreatedAt       time.Time       `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt       *gorm.DeletedAt `gorm:"index" json:"-"`
	ToNostr         bool            `gorm:"default:false" json:"toNostr"`
}
