package models

import "time"

type NoteRanking struct {
	NoteId      uint      `gorm:"primaryKey" json:"id"`
	Note        Note      `gorm:"foreignKey:NoteId" json:"note"`
	Reactions   int       `json:"reactions"`
	Replies     int       `json:"replies"`
	OnSatlantis bool      `json:"onSatlantis"`
	Nsfw        bool      `json:"nsfw"`
	Score       float64   `json:"score"`
	Timestamp   time.Time `json:"timestamp"`
}
