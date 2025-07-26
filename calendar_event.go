package models

import (
	"time"

	"gorm.io/gorm"
)

type CalendarEventType struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CalendarEvent struct {
	ID                 uint                        `gorm:"primaryKey" json:"id"`
	AccountID          uint                        `gorm:"index" json:"accountId"`
	Account            AccountDTO                  `json:"account"`
	Announcements      []CalendarEventAnnouncement `gorm:"foreignKey:CalendarEventID" json:"announcements"`
	ATag               string                      `json:"atag"`
	CreatedAt          *time.Time                  `json:"createdAt"`
	CalendarEventRSVPs []CalendarEventRSVP         `json:"calendarEventRsvps"`
	Content            *string                     `gorm:"type:text" json:"content"`
	DTag               string                      `json:"dtag"`
	End                time.Time                   `json:"end"`
	EndTzId            string                      `json:"endTzId"`
	EventID            uint                        `gorm:"index" json:"eventId"`
	Event              Event                       `json:"event"`
	Geohash            string                      `json:"geohash"`
	Image              string                      `json:"image"`
	Interests          []Interest                  `gorm:"many2many:calendar_event_interests" json:"interests"`
	Location           string                      `json:"location"`
	Kind               uint                        `gorm:"index" json:"kind"`
	NostrID            string                      `gorm:"index" json:"nostrId"`
	Notes              []CalendarEventNote         `gorm:"foreignKey:CalendarEventID" json:"notes"`
	PubKey             string                      `gorm:"type:text;index" json:"pubkey"`
	Sig                string                      `gorm:"type:text" json:"sig"`
	Start              time.Time                   `json:"start"`
	StartTzId          string                      `json:"startTzId"`
	Summary            string                      `json:"summary"`
	Tags               string                      `gorm:"type:jsonb" json:"tags"`
	Title              string                      `json:"title"`
	TypeID             uint                        `gorm:"index;not null;default:1" json:"typeId"`
	Type               CalendarEventType           `gorm:"foreignKey:TypeID" json:"type"`
	URL                string                      `json:"url"`
	Website            string                      `json:"website"`
	IsSatlantisCreated bool                        `gorm:"default:false" json:"isSatlantisCreated"`
	GoogleID           string                      `json:"googleId"`
	VenueID            *uint                       `gorm:"index" json:"venueId"`
	Venue              LocationDTO                 `gorm:"foreignKey:VenueID" json:"venue"`
	Cohosts            []CalendarEventCohost       `json:"cohosts"`
	DeletedAt          gorm.DeletedAt              `gorm:"index" json:"-"`
	OwnershipChangedAt *time.Time                  `json:"ownershipChangedAt"`
}

type CalendarEventInterest struct {
	CalendarEventID uint `gorm:"uniqueIndex:idx_calendar_event_interest"`
	InterestID      uint `gorm:"uniqueIndex:idx_calendar_event_interest"`
}

type CalendarEventCohost struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	CalendarEventID      uint           `gorm:"index;not null" json:"calendarEventId"`
	CalendarEvent        *CalendarEvent `gorm:"foreignKey:CalendarEventID" json:"calendarEvent,omitempty"`
	AccountID            uint           `gorm:"index;not null" json:"accountId"`
	Account              *AccountDTO    `gorm:"foreignKey:AccountID" json:"account"`
	InvitationAcceptedAt *time.Time     `json:"invitationAcceptedAt"`
	InvitationDeclinedAt *time.Time     `json:"invitationDeclinedAt"`
	AutoAcceptInvitation bool           `gorm:"default:false" json:"autoAcceptInvitation"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
}
