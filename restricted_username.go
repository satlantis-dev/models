package models

import (
	"time"
)

type RestrictedUsername struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"type:text" json:"username"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-,omitempty"`
}
