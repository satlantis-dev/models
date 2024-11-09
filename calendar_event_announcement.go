package models

type CalendarEventAnnouncement struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	CalendarEventID uint `json:"calendarEventId"`
	NoteID          uint `json:"noteId"`
	Note            Note `json:"note"`
}
