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
	ID                    uint                        `gorm:"primaryKey" json:"id"`
	AccountID             uint                        `gorm:"index" json:"accountId"`
	Account               *AccountDTO                 `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	EventID               uint                        `gorm:"index" json:"eventId"`
	NostrID               string                      `gorm:"index" json:"nostrId"`
	Event                 Event                       `json:"event"`
	Kind                  uint                        `gorm:"index" json:"kind"`
	Content               *string                     `gorm:"type:text" json:"content"`
	Tags                  string                      `gorm:"type:jsonb" json:"tags"`
	ATag                  string                      `json:"atag"`
	DTag                  string                      `json:"dtag"`
	Sig                   string                      `gorm:"type:text" json:"sig"`
	CreatedAt             *time.Time                  `json:"createdAt"`
	DeletedAt             gorm.DeletedAt              `gorm:"index" json:"-"`
	Announcements         []CalendarEventAnnouncement `gorm:"foreignKey:CalendarEventID" json:"announcements"`
	CalendarEventRSVPs    []CalendarEventRSVP         `json:"calendarEventRsvps"`
	Cohosts               []CalendarEventCohost       `json:"cohosts"`
	End                   time.Time                   `json:"end"`
	EndTzId               string                      `json:"endTzId"`
	Geohash               string                      `json:"geohash"`
	GoogleID              string                      `json:"googleId"`
	Image                 string                      `json:"image"`
	Interests             []Interest                  `gorm:"many2many:calendar_event_interests" json:"interests"`
	IsSatlantisCreated    bool                        `gorm:"default:false" json:"isSatlantisCreated"`
	Location              string                      `json:"location"`
	Notes                 []CalendarEventNote         `gorm:"foreignKey:CalendarEventID" json:"notes"`
	OwnershipChangedAt    *time.Time                  `json:"ownershipChangedAt"`
	PubKey                string                      `gorm:"type:text;index" json:"pubkey"`
	RsvpLimit             *int64                      `json:"rsvpLimit"`
	RsvpWaitlistEnabledAt *time.Time                  `json:"rsvpWaitlistEnabledAt"`
	Start                 time.Time                   `json:"start"`
	StartTzId             string                      `json:"startTzId"`
	Summary               string                      `json:"summary"`
	Title                 string                      `json:"title"`
	TypeID                uint                        `gorm:"index;not null;default:1" json:"typeId"`
	Type                  *CalendarEventType          `gorm:"foreignKey:TypeID" json:"type,omitempty"`
	URL                   string                      `json:"url"`
	Venue                 *LocationDTO                `gorm:"foreignKey:GoogleID;references:GoogleID;constraint:false" json:"venue,omitempty"`
	Website               string                      `json:"website"`
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
