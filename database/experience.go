package database

import (
	"time"
)

type Experience struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `gorm:"index" json:"-,omitempty"`
	AccountID   uint       `gorm:"index" json:"accountId"`
	CityID      uint       `gorm:"index" json:"cityId"`
	Cost        float32    `json:"cost"`
	Currency    string     `gorm:"type:text" json:"currency"`
	Description string     `gorm:"type:text" json:"description"`
	Name        string     `gorm:"type:text" json:"name"`
	Type        string     `gorm:"type:text" json:"type"`
	URL         string     `gorm:"type:text" json:"url"`
}
