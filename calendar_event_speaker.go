package models

import "time"

type CalendarEventSpeaker struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID      uint           `gorm:"not null;index;uniqueIndex:idx_speaker_event_account" json:"calendarEventId"`
	CalendarEvent        *CalendarEvent `gorm:"foreignKey:CalendarEventID" json:"calendarEvent,omitempty"`
	AccountID            uint           `gorm:"not null;index;uniqueIndex:idx_speaker_event_account" json:"accountId"`
	Account              *AccountDTO    `gorm:"foreignKey:AccountID" json:"account"`
	InvitationReceivedAt *time.Time     `gorm:"type:timestamptz" json:"invitationReceivedAt,omitempty"`
	InvitationAcceptedAt *time.Time     `gorm:"type:timestamptz" json:"invitationAcceptedAt,omitempty"`
	InvitationDeclinedAt *time.Time     `gorm:"type:timestamptz" json:"invitationDeclinedAt,omitempty"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
}
