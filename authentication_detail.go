package models

import (
	"time"
)

type AuthenticationProvider int

const (
	EmailProvider AuthenticationProvider = iota + 1
	GoogleProvider
	TwitterProvider
	FacebookProvider
	NostrProvider
)

type AuthenticationDetail struct {
	ID          uint                   `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time              `json:"-"`
	UpdatedAt   time.Time              `json:"-"`
	DeletedAt   *time.Time             `gorm:"index" json:"-,omitempty"`
	AccountID   uint                   `gorm:"index" json:"account_id"`
	Provider    AuthenticationProvider `gorm:"type:smallint" json:"provider"` // Email, Google, Twitter, Facebook, Nostr
	ProviderUID string                 `gorm:"unique_index" json:"-"`         // UID returned by the provider. For EmailProvider, it can be the same as the Account's email.
	Password    string                 `json:"password"`                      // Store encrypted password for EmailProvider. Not needed for social logins.
}
