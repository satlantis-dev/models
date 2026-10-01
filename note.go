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
	CalendarEventNoteDEPRECATED
	CalendarNote
	Ping
	ReactionNote
	DeleteNote
	ReplyNote
	MediaNote
)

type Note struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	AccountID          uint       `gorm:"index:idx_notes_account_type_created_at,priority:1;index:notes_created_on_satlantis_type_account_id_created_at_idx,priority:3" json:"accountId"`
	Account            AccountDTO `gorm:"-:migration" json:"account"`
	CreatedAt          *time.Time `gorm:"index:idx_notes_satlantis_created_at,priority:2;index:idx_notes_account_type_created_at,priority:3,sort:desc;index:notes_created_on_satlantis_type_account_id_created_at_idx,priority:4,sort:desc" json:"createdAt"`
	Content            *string    `gorm:"type:text" json:"content"`
	EventID            uint       `gorm:"unique" json:"eventId"`
	Event              *Event     `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"-"`
	Kind               uint       `gorm:"index" json:"kind"`
	NostrID            string     `gorm:"index" json:"nostrId"`
	PubKey             string     `gorm:"type:text;index" json:"pubkey"`
	Sig                string     `gorm:"type:text" json:"sig"`
	Tags               *string    `gorm:"type:jsonb" json:"tags"`
	Type               NoteType   `gorm:"index:idx_notes_account_type_created_at,priority:2;index:notes_created_on_satlantis_type_account_id_created_at_idx,priority:2" json:"type"`
	RepostedNoteID     *uint      `gorm:"index" json:"repostedNoteId"`
	RepostedNote       *Note      `json:"reposted_note" swaggerignore:"true"`
	CreatedOnSatlantis bool       `gorm:"index:idx_notes_satlantis_created_at,priority:1;index:notes_created_on_satlantis_type_account_id_created_at_idx,priority:1" json:"createdOnSatlantis"`
}

type NoteWithClosure struct {
	Note
	AncestorID   uint `gorm:"column:ancestor_id" json:"ancestorId"`
	Depth        int  `gorm:"column:depth" json:"depth"`
	DescendantID uint `gorm:"column:descendant_id" json:"descendantId"`
}

type FeedNote struct {
	Note
	Source               string           `json:"source"`
	Score                float64          `json:"score"`
	CommentCount         int              `json:"commentCount"`
	AllCommentCount      int              `json:"allCommentCount"`
	ReactionCount        int              `json:"reactionCount"`
	NotableReactionCount int              `json:"notableReactionCount"`
	ReactedByAccounts    []AccountMiniDTO `json:"reactedByAccounts"`
	RepostCount          int              `json:"repostCount"`
	RepostedByAccounts   []AccountMiniDTO `json:"repostedByAccounts"`
	CommentedByUser      bool             `json:"commentedByUser"`
	ReactedByUser        bool             `json:"reactedByUser"`
	Place                *PlaceDTO        `json:"place"`
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
