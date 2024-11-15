package models

import (
	"time"
)

type TeamMember struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-,omitempty"`
	AccountID uint       `gorm:"index" json:"accountId"`
	Account   AccountDTO `json:"account"`
	Bio       string     `json:"bio"`
	Name      string     `json:"name"`
	Npub      int64      `json:"npub"`
	Photo     string     `json:"photo"`
	Role      string     `json:"role"`
}
