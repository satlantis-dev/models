package models

import (
	"time"
)

type LocationNote struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `gorm:"index" json:"-,omitempty"`
	LocationID uint       `gorm:"index" json:"locationId"`
	Location   *Location  `gorm:"constraint:OnDelete:CASCADE;" json:"location,omitempty"`
	NoteID     uint       `gorm:"index" json:"noteId"`
	Note       *Note      `gorm:"constraint:OnDelete:CASCADE;" json:"note,omitempty"`
	Type       NoteType   `json:"type"`
}

type LocationFeedNote struct {
	ID         uint     `json:"id"`
	LocationID uint     `json:"locationId"`
	NoteID     uint     `json:"noteId"`
	Note       FeedNote `json:"note"`
	Type       NoteType `json:"type"`
}
