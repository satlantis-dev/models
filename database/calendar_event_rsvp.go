package database

type CalendarEventRSVP struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	AccountID uint    `gorm:"index;not null" json:"accountId"`
	Account   Account `json:"account"`
	EventID   uint    `gorm:"index;not null" json:"eventId"`
	NoteID    uint    `gorm:"index;not null" json:"noteId"`
	Status    string  `json:"status"`
}
