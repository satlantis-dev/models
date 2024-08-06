package database

import (
	"time"
)

type Continent struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-,omitempty"`
	Code      string     `gorm:"type:char(2);uniqueIndex" json:"code"`
	Name      string     `gorm:"type:text" json:"name"`
}
