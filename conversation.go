package models

import "time"

type Conversation struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	FromNpub          string     `gorm:"uniqueIndex:idx_from_npub_to_npub" json:"fromNpub"`
	ToNpub            string     `gorm:"uniqueIndex:idx_from_npub_to_npub" json:"toNpub"`
	FromAccountID     string     `gorm:"uniqueIndex:idx_from_account_to_account" json:"fromAccountID"`
	ToAccountID       string     `gorm:"uniqueIndex:idx_from_account_to_account" json:"toAccountID"`
	FromAccountReadAt *time.Time `json:"fromAccountReadAt"`
	ToAccountReadAt   *time.Time `json:"toAccountReadAt"`
	CreatedAt         time.Time  `json:"-"`
	DeletedAt         *time.Time `json:"-,omitempty"`
}
