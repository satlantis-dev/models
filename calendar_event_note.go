package models

type CalendarEventNote struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	CalendarEventID uint `json:"calendarEventId"`
	NoteID          uint `json:"noteId"`
	Note            Note `gorm:"constraint:OnDelete:CASCADE;" json:"note"`
}
