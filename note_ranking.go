package models

import "time"

type NoteRanking struct {
	NoteId           uint      `gorm:"primaryKey" json:"noteId"`
	Note             Note      `gorm:"foreignKey:NoteId" json:"note"`
	OnSatlantis      bool      `json:"onSatlantis"`
	FromFocusAccount bool      `json:"fromFocusAccount"`
	Nsfw             bool      `json:"nsfw"`
	Reactions        int       `gorm:"default:0" json:"reactions"`
	Replies          int       `gorm:"default:0" json:"replies"`
	AllReplies       int       `gorm:"default:0" json:"allReplies"`
	Score            float64   `gorm:"default:0" json:"score"`
	Timestamp        time.Time `json:"timestamp"`
}
