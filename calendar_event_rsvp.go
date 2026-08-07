package models

import "time"

type RsvpStatus string

var (
	// RsvpAccepted means the guest is confirmed to attend - set by the guest
	// accepting (directly, or by accepting an invite), or by an organizer/
	// admin approving a request or assigning a ticket.
	RsvpAccepted RsvpStatus = "accepted"
	// RsvpTentative means the guest might attend - set by the guest.
	RsvpTentative RsvpStatus = "tentative"
	// RsvpDeclined means the guest said no - set by the guest, either
	// declining an invitation (Invited -> Declined) or reversing their own
	// earlier acceptance (Accepted -> Declined).
	RsvpDeclined RsvpStatus = "declined"
	// RsvpWaitlisted means the guest is confirmed but
	// the event was full when they were accepted (either accepting directly
	// or via organizer approval of a gated request), so they're first in
	// line to be promoted to Accepted once a slot frees up.
	RsvpWaitlisted RsvpStatus = "waitlisted"
	// RsvpRequested means the guest is awaiting organizer/admin review - the
	// initial status for any new RSVP on a gated event.
	RsvpRequested RsvpStatus = "requested"
	// RsvpRejected means an organizer/admin said no - either denying a
	// Requested (gated) submission before it was ever accepted, or actively
	// revoking an existing RSVP.
	RsvpRejected RsvpStatus = "rejected"
	// RsvpInvited means an organizer/admin sent this account an invitation
	// that hasn't been responded to yet.
	RsvpInvited RsvpStatus = "invited"
)

type CalendarEventRSVP struct {
	ID                  uint                        `gorm:"primaryKey" json:"id"`
	AccountID           uint                        `gorm:"index:idx_rsvp_account_event,priority:1" json:"accountId"`
	Account             AccountDTO                  `gorm:"constraint:OnDelete:CASCADE;" json:"account"`
	CreatedAt           time.Time                   `json:"createdAt"`
	CalendarEventID     uint                        `gorm:"index:idx_rsvp_account_event,priority:2" json:"calendarEventId"`
	Status              string                      `json:"status"`
	AcceptedAt          *time.Time                  `json:"acceptedAt,omitempty"`
	RejectedAt          *time.Time                  `json:"rejectedAt,omitempty"`
	StatusUpdatedById   *uint                       `json:"-"`
	StatusUpdatedBy     *Account                    `json:"-"`
	IsSatlantisCreated  bool                        `json:"-" gorm:"default:false"`
	RegistrationAnswers *RegistrationAnswersPayload `gorm:"type:jsonb;serializer:json" json:"registrationAnswers,omitempty"`
	InvitationMessage   *string                     `gorm:"type:text" json:"invitationMessage,omitempty"`

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
