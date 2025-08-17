package models

import "time"

type RsvpStatus string

var (
	RsvpAccepted   RsvpStatus = "accepted"
	RsvpTentative  RsvpStatus = "tentative"
	RsvpDeclined   RsvpStatus = "declined"
	RsvpWaitlisted RsvpStatus = "waitlisted"
	RsvpPending    RsvpStatus = "pending"
)

type CalendarEventRSVP struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	AccountID            uint      `json:"accountId"`
	Account              Account   `json:"account"`
	CreatedAt            time.Time `json:"createdAt"`
	Content              string    `gorm:"type:text" json:"content"`
	EventID              uint      `gorm:"index" json:"eventId"`
	Kind                 uint      `gorm:"index" json:"kind"`
	NostrID              string    `gorm:"index" json:"nostrId"`
	PubKey               string    `gorm:"type:text;index" json:"pubkey"`
	Sig                  string    `gorm:"type:text" json:"sig"`
	Tags                 string    `gorm:"type:jsonb" json:"tags"`
	CalendarEventID      uint      `json:"calendarEventId"`
	Status               string    `json:"status"`
	NotificationWeekSent bool      `gorm:"default:false"`
	NotificationDaySent  bool      `gorm:"default:false"`
	NotificationHourSent bool      `gorm:"default:false"`
	IsSatlantisCreated   bool      `gorm:"default:false"`
}
