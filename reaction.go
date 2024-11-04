package models

type Reaction struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	AccountID uint    `gorm:"index;not null" json:"accountId"`
	Account   Account `json:"account"`
	EventID   uint    `gorm:"index;not null" json:"eventId"`
	Event     Event   `json:"event"`
	NoteID    uint    `gorm:"index;not null" json:"noteId"`
}
