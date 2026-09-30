package models

import (
	"time"
)

type Continent struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Code      string     `gorm:"type:bpchar(2);uniqueIndex" json:"code"`
	Name      string     `gorm:"type:text" json:"name"`
}
