package models

import "time"

type RsvpStatus string

var (
	RsvpAccepted   RsvpStatus = "accepted"
	RsvpTentative  RsvpStatus = "tentative"
	RsvpDeclined   RsvpStatus = "declined"
	RsvpWaitlisted RsvpStatus = "waitlisted"
	RsvpPending    RsvpStatus = "pending"
	RsvpRejected   RsvpStatus = "rejected"
)

type CalendarEventRSVP struct {
	ID                  uint       `gorm:"primaryKey" json:"id"`
	AccountID           uint       `json:"accountId"`
	Account             Account    `json:"account"`
	CreatedAt           time.Time  `json:"createdAt"`
	CalendarEventID     uint       `json:"calendarEventId"`
	Status              string     `json:"status"`
	AcceptedAt          *time.Time `json:"acceptedAt,omitempty"`
	RejectedAt          *time.Time `json:"rejectedAt,omitempty"`
	StatusUpdatedById   *uint      `json:"statusUpdatedById"`
	StatusUpdatedBy     *Account   `json:"statusUpdatedBy"`
	IsSatlantisCreated  bool       `gorm:"default:false"`
	RegistrationAnswers *string    `gorm:"type:jsonb" json:"registrationAnswers"`

	// Nostr fields
	Content string `gorm:"type:text" json:"content"`
	EventID uint   `gorm:"index" json:"eventId"`
	Kind    uint   `gorm:"index" json:"kind"`
	NostrID string `gorm:"index" json:"nostrId"`
	PubKey  string `gorm:"type:text;index" json:"pubkey"`
	Sig     string `gorm:"type:text" json:"sig"`
	Tags    string `gorm:"type:jsonb" json:"tags"`

	// Notification flags
	NotificationWeekSent bool `gorm:"default:false"`
	NotificationDaySent  bool `gorm:"default:false"`
	NotificationHourSent bool `gorm:"default:false"`
}
