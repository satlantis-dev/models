package models

import "time"

type RsvpStatus string

var (
	RsvpAccepted   RsvpStatus = "accepted"
	RsvpTentative  RsvpStatus = "tentative"
	RsvpDeclined   RsvpStatus = "declined"
	RsvpWaitlisted RsvpStatus = "waitlisted"
	RsvpRequested  RsvpStatus = "requested"
	RsvpRejected   RsvpStatus = "rejected"
	RsvpInvited    RsvpStatus = "invited"
)

type CalendarEventRSVP struct {
	ID                  uint                    `gorm:"primaryKey" json:"id"`
	AccountID           uint                    `json:"accountId"`
	Account             AccountDTO              `gorm:"constraint:OnDelete:CASCADE;" json:"account"`
	CreatedAt           time.Time               `json:"createdAt"`
	CalendarEventID     uint                    `json:"calendarEventId"`
	Status              string                  `json:"status"`
	AcceptedAt          *time.Time              `json:"acceptedAt,omitempty"`
	RejectedAt          *time.Time              `json:"rejectedAt,omitempty"`
	StatusUpdatedById   *uint                   `json:"-"`
	StatusUpdatedBy     *Account                `json:"-"`
	IsSatlantisCreated  bool                    `json:"-" gorm:"default:false"`
	RegistrationAnswers *map[string]interface{} `gorm:"type:jsonb;serializer:json" json:"registrationAnswers,omitempty"`
	InvitationMessage   *string                 `gorm:"type:text" json:"invitationMessage,omitempty"`

	// Nostr fields
	Content string `gorm:"type:text" json:"-"`
	EventID uint   `gorm:"index" json:"-"`
	Kind    uint   `gorm:"index" json:"-"`
	NostrID string `gorm:"index" json:"-"`
	PubKey  string `gorm:"type:text;index" json:"-"`
	Sig     string `gorm:"type:text" json:"-"`
	Tags    string `gorm:"type:jsonb" json:"-"`

	// Notification flags
	NotificationHourSentAt *time.Time `json:"-"`
	NotificationDaySentAt  *time.Time `json:"-"`
	NotificationWeekSentAt *time.Time `json:"-"`
}
