package models

import (
	"time"
)

type NoteType int

const (
	BasicNote NoteType = iota + 1
	ReviewNote
	GalleryNote
	PublicChatNote
	PrivateChatNote
	CalendarEventNote
	CalendarNote
	Ping
	ReactionNote
	DeleteNote
	ReplyNote
	MediaNote
)

type Note struct {
	ID                 uint                `gorm:"primaryKey" json:"id"`
	AccountID          uint                `gorm:"index" json:"accountId"`
	Account            Account             `json:"account"`
	CreatedAt          *time.Time          `json:"createdAt"`
	CalendarEventRSVPs []CalendarEventRSVP `gorm:"foreignKey:NoteID" json:"calendarEventRsvps"`
	ChatMemberships    []ChatMembership    `gorm:"foreignKey:NoteID" json:"chatMemberships"`
	Content            *string             `gorm:"type:text" json:"content"`
	Descendants        []NoteWithClosure   `gorm:"-" json:"descendants"`
	EventID            uint                `gorm:"index" json:"eventId"`
	Event              Event               `json:"event"`
	Kind               uint                `gorm:"index" json:"kind"`
	NostrID            string              `gorm:"index" json:"nostrId"`
	PubKey             string              `gorm:"type:text;index" json:"pubkey"`
	Sig                string              `gorm:"type:text" json:"sig"`
	Tags               string              `gorm:"type:jsonb" json:"tags"`
	Type               NoteType            `json:"type"`
	RepostedNoteID     *uint               `gorm:"index" json:"repostedNoteId"`
	RepostedNote       *Note               `json:"reposted_note"`
	Reactions          []Reaction          `gorm:"foreignKey:NoteID" json:"reactions"`
}

type NoteWithClosure struct {
	Note
	AncestorID   uint `gorm:"column:ancestor_id" json:"ancestorId"`
	Depth        int  `gorm:"column:depth" json:"depth"`
	DescendantID uint `gorm:"column:descendant_id" json:"descendantId"`
}

type ChatNote struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	AccountID uint       `gorm:"index" json:"accountId"`
	Account   Account    `json:"account" gorm:"foreignKey:AccountID"`
	EventID   uint       `gorm:"index" json:"eventId"`
	Event     Event      `json:"event"`
	Reactions []Reaction `gorm:"foreignKey:NoteID" json:"reactions"`
}

type NoteWithStartTime struct {
	Note      NoteWithClosure
	StartTime time.Time
}

type NotePagination struct {
	PaginationForward bool `json:"paginationForward"`
	PaginationLimit   int  `json:"paginationLimit"`
	PaginationNoteID  int  `json:"paginationNoteId"`
}
