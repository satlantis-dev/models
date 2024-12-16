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
	ID             uint       `gorm:"primaryKey" json:"id"`
	AccountID      uint       `gorm:"index" json:"accountId"`
	Account        AccountDTO `json:"account"`
	CreatedAt      *time.Time `json:"createdAt"`
	Content        *string    `gorm:"type:text" json:"content"`
	EventID        uint       `gorm:"index;unique" json:"eventId"`
	Kind           uint       `gorm:"index" json:"kind"`
	NostrID        string     `gorm:"index" json:"nostrId"`
	PubKey         string     `gorm:"type:text;index" json:"pubkey"`
	Sig            string     `gorm:"type:text" json:"sig"`
	Tags           *string    `gorm:"type:jsonb" json:"tags"`
	Type           NoteType   `json:"type"`
	RepostedNoteID *uint      `gorm:"index" json:"repostedNoteId"`
	RepostedNote   *Note      `json:"reposted_note"`
}

type NoteWithClosure struct {
	Note
	AncestorID   uint `gorm:"column:ancestor_id" json:"ancestorId"`
	Depth        int  `gorm:"column:depth" json:"depth"`
	DescendantID uint `gorm:"column:descendant_id" json:"descendantId"`
}

type FeedNote struct {
	Note
	Source          string  `json:"source"`
	Score           float64 `json:"score"`
	CommentCount    int     `json:"commentCount"`
	AllCommentCount int     `json:"allCommentCount"`
	ReactionCount   int     `json:"reactionCount"`
	CommentedByUser bool    `json:"commentedByUser"`
	ReactedByUser   bool    `json:"reactedByUser"`
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

type NoteDTO struct {
	ID        uint       `json:"id"`
	AccountID uint       `json:"accountId"`
	Account   AccountDTO `json:"account"`
	CreatedAt *time.Time `json:"createdAt"`
	Content   *string    `json:"content"`
	EventID   uint       `json:"eventId"`
	Kind      uint       `json:"kind"`
	NostrID   string     `json:"nostrId"`
	PubKey    string     `json:"pubkey"`
	Sig       string     `json:"sig"`
	Tags      *string    `json:"tags"`
	Type      NoteType   `json:"type"`
	Place     *Place     `json:"place"`
}

func (note *Note) ToDTO() NoteDTO {
	return NoteDTO{
		ID:        note.ID,
		AccountID: note.AccountID,
		Account:   note.Account,
		CreatedAt: note.CreatedAt,
		Content:   note.Content,
		EventID:   note.EventID,
		Kind:      note.Kind,
		NostrID:   note.NostrID,
		PubKey:    note.PubKey,
		Sig:       note.Sig,
		Tags:      note.Tags,
		Type:      note.Type,
	}
}
