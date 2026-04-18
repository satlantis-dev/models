package models

type CalendarCalendarEvent struct {
	CalendarID      uint           `gorm:"index;uniqueIndex:idx_calendar_calendar_event" json:"calendarId"`
	Calendar        *Calendar      `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE" json:"calendar,omitempty"`
	CalendarEventID uint           `gorm:"index;uniqueIndex:idx_calendar_calendar_event" json:"calendarEventId"`
	CalendarEvent   *CalendarEvent `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE" json:"calendarEvent,omitempty"`
	Featured        bool           `gorm:"default:false" json:"featured"`
}

func (CalendarCalendarEvent) TableName() string {
	return "calendar_calendar_events"
}
