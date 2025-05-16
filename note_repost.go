package models

import "time"

type NoteRepost struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	EventID   string     `gorm:"index" json:"eventId"`
	NoteID    uint       `json:"noteId"`
	Note      Note       `gorm:"foreignKey:NoteID" json:"note"`
	AccountID uint       `json:"accountId"`
	Account   Account    `gorm:"foreignKey:AccountID" json:"account"`
	CreatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
