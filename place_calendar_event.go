package models

import "time"

type PlaceCalendarEvent struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       *time.Time     `gorm:"index" json:"-,omitempty"`
	PlaceID         uint           `gorm:"index" json:"placeId"`
	Place           *Place         `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE;" json:"place,omitempty"`
	CalendarEventID uint           `gorm:"index" json:"calendarEventId"`
	CalendarEvent   *CalendarEvent `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE;" json:"calendarEvent,omitempty"`
}
