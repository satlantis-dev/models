package models

import "time"

type PlaceCalendarEvent struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time     `json:"-"`
	UpdatedAt       time.Time     `json:"-"`
	DeletedAt       *time.Time    `gorm:"index" json:"-,omitempty"`
	PlaceID         uint          `gorm:"index" json:"placeId"`
	CalendarEventID uint          `gorm:"index" json:"calendarEventId"`
	CalendarEvent   CalendarEvent `json:"calendarEvent"`
}
