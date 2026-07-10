package models

import "github.com/lib/pq"

type CalendarCalendarEvent struct {
	CalendarID      uint           `gorm:"primaryKey;index;uniqueIndex:idx_calendar_calendar_event" json:"calendarId"`
	Calendar        *Calendar      `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE" json:"calendar,omitempty"`
	CalendarEventID uint           `gorm:"primaryKey;index;uniqueIndex:idx_calendar_calendar_event" json:"calendarEventId"`
	CalendarEvent   *CalendarEvent `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"calendarEvent,omitempty"`
	Featured        bool           `gorm:"default:false" json:"featured"`

	// MembersOnly restricts this event's visibility, within this calendar, to
	// community members. Defaults to false (visible to non-members too).
	MembersOnly bool `gorm:"not null;default:false" json:"membersOnly"`

	// TierIDsOnly further limits visibility, when MembersOnly is true, to
	// members of the listed community membership tiers (by ID). Empty/nil means
	// the event is visible to members of any tier.
	TierIDsOnly pq.Int32Array `gorm:"type:integer[]" json:"tierIdsOnly,omitempty"`
}

func (CalendarCalendarEvent) TableName() string {
	return "calendar_calendar_events"
}
