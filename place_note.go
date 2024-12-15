package models

import (
	"time"
)

type PlaceNote struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-,omitempty"`
	PlaceID   uint       `gorm:"index" json:"placeId"`
	NoteID    uint       `gorm:"index" json:"noteId"`
	Note      Note       `json:"note"`
	Nsfw      bool       `json:"nsfw"`
	Type      NoteType   `json:"type"`
}

type PlaceFeedNote struct {
	ID      uint     `json:"id"`
	PlaceID uint     `json:"placeId"`
	NoteID  uint     `json:"noteId"`
	Note    FeedNote `json:"note"`
	Type    NoteType `json:"type"`
}

type PlaceNoteWithClosure struct {
	PlaceNote
	AncestorID   uint `gorm:"column:ancestor_id" json:"ancestorId"`
	Depth        int  `gorm:"column:depth" json:"depth"`
	DescendantID uint `gorm:"column:descendant_id" json:"descendantId"`
}

type PlaceNoteWithStartTime struct {
	Note      PlaceNoteWithClosure
	StartTime time.Time
}
