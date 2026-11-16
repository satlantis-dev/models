package models

import "time"

// CalendarEventSpeakerNoticeMaxLength is the longest notice, in characters, a
// host can set on a speaker.
const CalendarEventSpeakerNoticeMaxLength = 250

type CalendarEventSpeaker struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID      uint           `gorm:"not null;uniqueIndex:idx_speaker_event_account" json:"calendarEventId"`
	CalendarEvent        *CalendarEvent `gorm:"foreignKey:CalendarEventID" json:"calendarEvent,omitempty"`
	AccountID            uint           `gorm:"not null;index;uniqueIndex:idx_speaker_event_account" json:"accountId"`
	Account              *AccountDTO    `gorm:"foreignKey:AccountID" json:"account"`
	InvitationReceivedAt *time.Time     `gorm:"type:timestamptz" json:"invitationReceivedAt,omitempty"`
	InvitationAcceptedAt *time.Time     `gorm:"type:timestamptz" json:"invitationAcceptedAt,omitempty"`
	InvitationDeclinedAt *time.Time     `gorm:"type:timestamptz" json:"invitationDeclinedAt,omitempty"`
	// Notice is the host's notice about the speaker for this event, at most
	// CalendarEventSpeakerNoticeMaxLength characters.
	Notice *string `gorm:"type:text" json:"notice"`
	// Visible shows the speaker in the event's public lineup. Only a validated
	// speaker (accepted, or added without an invitation) can be visible.
	Visible bool `gorm:"not null;default:false" json:"visible"`
	// Rank orders the event's lineup, lowest first (ties by CreatedAt).
	Rank      int       `gorm:"not null;default:0" json:"rank"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
