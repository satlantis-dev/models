package models

import "time"

type NoteRepost struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	EventID   string     `gorm:"index" json:"eventId"`
	NoteID    uint       `gorm:"index:,option:CONCURRENTLY" json:"noteId"`
	Note      Note       `gorm:"foreignKey:NoteID;constraint:OnDelete:CASCADE" json:"note"`
	AccountID uint       `gorm:"index:,option:CONCURRENTLY" json:"accountId"`
	Account   Account    `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account"`
	CreatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
