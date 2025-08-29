package models

import "time"

type PlaceCalendarEvent struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	UpdatedAt       time.Time     `json:"-"`
	DeletedAt       *time.Time    `gorm:"index" json:"-,omitempty"`
	PlaceID         uint          `gorm:"index;constraint:OnDelete:CASCADE;" json:"placeId"`
	Place           Place         `json:"place,omitempty"`
	CalendarEventID uint          `gorm:"index;constraint:OnDelete:CASCADE;" json:"calendarEventId"`
	CalendarEvent   CalendarEvent `json:"calendarEvent,omitempty"`
}
