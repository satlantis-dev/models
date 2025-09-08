package models

type CalendarEventAnnouncement struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID uint           `json:"calendarEventId"`
	CalendarEvent   *CalendarEvent `gorm:"constraint:OnDelete:CASCADE;" json:"calendarEvent,omitempty"`
	NoteID          uint           `json:"noteId"`
	Note            *Note          `gorm:"constraint:OnDelete:CASCADE;" json:"note,omitempty"`
}
