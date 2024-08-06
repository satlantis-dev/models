package models

import (
	"time"

	"github.com/nbd-wtf/go-nostr"
)

type ChatMembership struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `gorm:"index" json:"-,omitempty"`
	AccountID      uint       `gorm:"index" json:"accountId"`
	Account        Account    `json:"account"`
	LastReadNoteID *uint      `json:"lastReadNoteId"`
	NoteID         uint       `gorm:"index" json:"noteId"`
	Note           Note       `json:"note"`
}

type PublishNote struct {
	ChatNoteID     uint         `json:"noteId"`
	AccountID      uint         `json:"accountId"`
	Event          *nostr.Event `json:"event"`
	Type           NoteType     `json:"type"`
	ParentID       uint         `json:"parentId"`
	RepostedNoteID *uint        `json:"repostedNoteId"`
	PlaceID        *uint        `json:"placeId"`
}

type NoteResponse struct {
	ChatMembership ChatMembership `json:"chatMembership"`
	Note           Note           `json:"note"`
}

type ChatInfo struct {
	LastReadNoteID uint `json:"lastReadNoteId"`
	UnreadCount    uint `json:"unreadCount"`
	LastMessage    Note `json:"lastMessage"`
}
