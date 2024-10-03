package models

import "time"

type CalendarEventRSVP struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	AccountID       uint      `gorm:"index;not null" json:"accountId"`
	Account         Account   `json:"account"`
	CreatedAt       time.Time `json:"createdAt"`
	Content         string    `gorm:"type:text" json:"content"`
	EventID         uint      `gorm:"index" json:"eventId"`
	Event           Event     `json:"event"`
	NostrID         string    `gorm:"index" json:"nostrId"`
	PubKey          string    `gorm:"type:text;index" json:"pubkey"`
	Sig             string    `gorm:"type:text" json:"sig"`
	Tags            string    `gorm:"type:jsonb" json:"tags"`
	CalendarEventID uint      `gorm:"index;not null" json:"calendarEventId"`
	Status          string    `json:"status"`
}
