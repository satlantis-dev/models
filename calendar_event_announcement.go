package models

type CalendarEventAnnouncement struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	CalendarEventID uint    `json:"calendarEventId"`
	NoteID          uint    `json:"noteId"`
	Note            *Note   `gorm:"constraint:OnDelete:CASCADE;" json:"note,omitempty"`
	ToDiscussion    bool    `gorm:"default:true" json:"toDiscussion"`
	ToNostr         bool    `gorm:"default:false" json:"toNostr"`
	ToEmail         bool    `gorm:"default:false" json:"toEmail"`
	EmailSubject    *string `gorm:"type:text" json:"emailSubject,omitempty"`
}
