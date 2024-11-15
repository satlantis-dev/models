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
	NoteID     uint       `gorm:"index" json:"noteId"`
	Note       Note       `json:"note"`
	Type       NoteType   `json:"type"`
}

type LocationFeedNote struct {
	ID         uint     `json:"id"`
	LocationID uint     `json:"locationId"`
	NoteID     uint     `json:"noteId"`
	Note       FeedNote `json:"note"`
	Type       NoteType `json:"type"`
}

type LocationNoteWithClosure struct {
	LocationNote
	AncestorID   uint `gorm:"column:ancestor_id" json:"ancestorId"`
	Depth        int  `gorm:"column:depth" json:"depth"`
	DescendantID uint `gorm:"column:descendant_id" json:"descendantId"`
}

type LocationNoteWithStartTime struct {
	Note      LocationNoteWithClosure
	StartTime time.Time
}
