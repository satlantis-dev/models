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
	AccountID             uint                        `gorm:"index;not null" json:"accountId"`
	Account               *AccountDTO                 `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	ContactEmail          *string                     `gorm:"type:text" json:"contactEmail"`
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
	UpdatedAt             *time.Time                  `json:"updatedAt"`
	DeletedAt             gorm.DeletedAt              `gorm:"index" json:"-"`
	Announcements         []CalendarEventAnnouncement `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE;" json:"announcements"`
	CalendarEventRSVPs    []CalendarEventRSVP         `json:"calendarEventRsvps"`
	Cohosts               []CalendarEventCohost       `json:"cohosts"`
	End                   time.Time                   `json:"end"`
	EndTzId               string                      `gorm:"not null" json:"endTzId"`
	Featured              bool                        `gorm:"default:false" json:"featured"`
	Geohash               string                      `json:"geohash"`
	GoogleID              string                      `json:"googleId"`
	Image                 string                      `json:"image"`
	CalendarEventTags     []CalendarEventTag          `gorm:"many2many:calendar_event_calendar_event_tags;constraint:OnDelete:CASCADE" json:"calendarEventTags"`
	IsSatlantisCreated    bool                        `gorm:"default:false" json:"isSatlantisCreated"`
	IsUnlisted            bool                        `gorm:"default:false;index" json:"isUnlisted"`
	IsHidingAttendees     bool                        `gorm:"default:false" json:"isHidingAttendees"`
	Location              string                      `json:"location"`
	Notes                 []CalendarEventNote         `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:CASCADE;" json:"notes"`
	OwnershipChangedAt    *time.Time                  `json:"ownershipChangedAt"`
	PlaceID               *uint                       `gorm:"index" json:"placeId"`
	Place                 *PlaceDTO                   `gorm:"foreignKey:PlaceID" json:"place,omitempty"`
	PubKey                string                      `gorm:"type:text;index" json:"pubkey"`
	RsvpLimit             *int64                      `json:"rsvpLimit"`
	RsvpWaitlistEnabledAt *time.Time                  `json:"rsvpWaitlistEnabledAt"`
	RsvpGatedEnabledAt    *time.Time                  `json:"rsvpGatedEnabledAt"`
	Start                 time.Time                   `json:"start"`
	StartTzId             string                      `gorm:"not null" json:"startTzId"`
	Summary               string                      `json:"summary"`
	Title                 string                      `json:"title"`
	TypeID                uint                        `gorm:"index;not null;default:1" json:"typeId"`
	Type                  *CalendarEventType          `gorm:"foreignKey:TypeID" json:"type,omitempty"`
	URL                   string                      `json:"url"`
	Venue                 *LocationDTO                `gorm:"-" json:"venue,omitempty"`
	Website               string                      `json:"website"`
	RegistrationQuestions *map[string]interface{}     `gorm:"type:jsonb;serializer:json" json:"registrationQuestions,omitempty"`
	OfficialCalendarID    *uint                       `gorm:"index" json:"officialCalendarId,omitempty"`
	OfficialCalendar      *Calendar                   `gorm:"foreignKey:OfficialCalendarID" json:"officialCalendar,omitempty"`
}

type CalendarEventDTO struct {
	ID                 uint               `gorm:"primaryKey" json:"id"`
	AccountID          uint               `gorm:"index;not null" json:"accountId"`
	Account            *AccountDTO        `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	End                time.Time          `json:"end"`
	EndTzId            string             `gorm:"not null" json:"endTzId"`
	Featured           bool               `gorm:"default:false" json:"featured"`
	GoogleID           string             `json:"googleId"`
	Image              string             `json:"image"`
	IsUnlisted         bool               `gorm:"default:false;index" json:"isUnlisted"`
	IsHidingAttendees  bool               `gorm:"default:false" json:"isHidingAttendees"`
	Location           string             `json:"location"`
	PlaceID            *uint              `gorm:"index" json:"placeId"`
	Place              *PlaceDTO          `gorm:"foreignKey:PlaceID" json:"place,omitempty"`
	Start              time.Time          `json:"start"`
	StartTzId          string             `gorm:"not null" json:"startTzId"`
	Summary            string             `json:"summary"`
	Title              string             `json:"title"`
	TypeID             uint               `gorm:"index;not null;default:1" json:"typeId"`
	Type               *CalendarEventType `gorm:"foreignKey:TypeID" json:"type,omitempty"`
	URL                string             `json:"url"`
	Website            string             `json:"website"`
	OfficialCalendarID *uint              `gorm:"index" json:"officialCalendarId,omitempty"`
	OfficialCalendar   *Calendar          `gorm:"foreignKey:OfficialCalendarID" json:"officialCalendar,omitempty"`
}

func (CalendarEventDTO) TableName() string {
	return "calendar_events"
}

func (c CalendarEvent) ToDTO() *CalendarEventDTO {

	return &CalendarEventDTO{
		ID:                 c.ID,
		AccountID:          c.AccountID,
		Account:            c.Account,
		End:                c.End,
		EndTzId:            c.EndTzId,
		Featured:           c.Featured,
		GoogleID:           c.GoogleID,
		Image:              c.Image,
		IsUnlisted:         c.IsUnlisted,
		IsHidingAttendees:  c.IsHidingAttendees,
		Location:           c.Location,
		PlaceID:            c.PlaceID,
		Place:              c.Place,
		Start:              c.Start,
		StartTzId:          c.StartTzId,
		Summary:            c.Summary,
		Title:              c.Title,
		TypeID:             c.TypeID,
		Type:               c.Type,
		URL:                c.URL,
		Website:            c.Website,
		OfficialCalendarID: c.OfficialCalendarID,
		OfficialCalendar:   c.OfficialCalendar,
	}
}

type CalendarEventCalendarEventTag struct {
	CalendarEventID    uint `gorm:"uniqueIndex:idx_calendar_event_calendar_event_tag"`
	CalendarEventTagID uint `gorm:"uniqueIndex:idx_calendar_event_calendar_event_tag"`
}

type CalendarEventResponse struct {
	*CalendarEventDTO
	Country         map[string]interface{} `json:"country,omitempty"`
	Venue           *LocationDTO           `json:"venue,omitempty"`
	AttendeeCount   int                    `json:"attendeeCount"`
	SampleAttendees *[]AccountDTO          `json:"sampleAttendees,omitempty"`
	KnownAttendees  *[]AccountDTO          `json:"knownAttendees,omitempty"`
}

// UserRSVPInfo contains minimal RSVP information for the requesting user
type UserRSVPInfo struct {
	ID        uint       `json:"id"`
	AccountID uint       `json:"accountId"`
	Status    RsvpStatus `json:"status"`
}

// UserTicketInfo contains minimal ticket information for the requesting user
type UserTicketInfo struct {
	ID             uint         `json:"id"`
	AccountID      uint         `json:"accountId"`
	Status         TicketStatus `json:"status"`
	TicketTypeID   uint         `json:"ticketTypeId"`
	TicketTypeName string       `json:"ticketTypeName"`
	CreatedAt      time.Time    `json:"createdAt"`
	CheckedInAt    *time.Time   `json:"checkedInAt,omitempty"`
}

type CalendarEventResponseFull struct {
	*CalendarEvent
	RsvpAcceptedCount   int64           `json:"rsvpAcceptedCount"`
	RsvpWaitlistedCount int64           `json:"rsvpWaitlistedCount"`
	RsvpInvitedCount    int64           `json:"rsvpInvitedCount"`
	RsvpTentativeCount  int64           `json:"rsvpTentativeCount"`
	RsvpRequestedCount  int64           `json:"rsvpRequestedCount"`
	RsvpDeclinedCount   int64           `json:"rsvpDeclinedCount"`
	RsvpRejectedCount   int64           `json:"rsvpRejectedCount"`
	Venue               *LocationDTO    `json:"venue,omitempty"`
	UserRSVP            *UserRSVPInfo   `json:"userRsvp,omitempty"`
	UserTicket          *UserTicketInfo `json:"userTicket,omitempty"`
}
