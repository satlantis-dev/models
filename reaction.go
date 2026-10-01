package models

type Reaction struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	AccountID uint       `gorm:"index;not null" json:"accountId"`
	Account   AccountDTO `json:"account"`
	EventID   uint       `gorm:"index;not null" json:"eventId"`
	Event     Event      `gorm:"constraint:OnDelete:CASCADE" json:"event"`
	NoteID    uint       `gorm:"index;not null" json:"noteId"`
	Note      *Note      `gorm:"constraint:OnDelete:CASCADE;" json:"note,omitempty"`
}
