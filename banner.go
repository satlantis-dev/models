package models

import "time"

type Banner struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Url       string    `gorm:"not null;unique" json:"url"`
	Category  string    `gorm:"not null" json:"category"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
}
