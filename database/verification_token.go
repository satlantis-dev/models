package database

import (
	"time"
)

type VerificationToken struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-,omitempty"`
	Token     string     `gorm:"type:text" json:"token"`
	AccountID uint       `gorm:"index" json:"accountId"`
}
