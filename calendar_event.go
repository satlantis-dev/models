package models

import "time"

type CalendarEventType string

const (
	Conference CalendarEventType = "conference"
	Meetup     CalendarEventType = "meetup"
	Hackathon  CalendarEventType = "hackathon"
	Concert    CalendarEventType = "concert"
	Workshop   CalendarEventType = "workshop"
	Party      CalendarEventType = "party"
	Play       CalendarEventType = "play"
	Sports     CalendarEventType = "sports"
	Exhibition CalendarEventType = "exhibition"
	Festival   CalendarEventType = "festival"
	Music      CalendarEventType = "music"
	Other      CalendarEventType = "other"
)

type CalendarEvent struct {
	ID                 uint                `gorm:"primaryKey" json:"id"`
	AccountID          uint                `gorm:"index" json:"accountId"`
	Account            Account             `json:"account"`
	Announcements      []CalendarEventNote `gorm:"foreignKey:CalendarEventID" json:"announcements"`
	ATag               string              `json:"atag"`
	CreatedAt          *time.Time          `json:"createdAt"`
	CalendarEventRSVPs []CalendarEventRSVP `json:"calendarEventRsvps"`
	Content            *string             `gorm:"type:text" json:"content"`
	DTag               string              `json:"dtag"`
	End                time.Time           `json:"end"`
	EndTzId            string              `json:"endTzId"`
	EventID            uint                `gorm:"index" json:"eventId"`
	Event              Event               `json:"event"`
	Geohash            string              `json:"geohash"`
	Image              string              `json:"image"`
	Location           string              `json:"location"`
	Kind               uint                `gorm:"index" json:"kind"`
	NostrID            string              `gorm:"index" json:"nostrId"`
	Notes              []CalendarEventNote `gorm:"foreignKey:CalendarEventID" json:"notes"`
	PubKey             string              `gorm:"type:text;index" json:"pubkey"`
	Sig                string              `gorm:"type:text" json:"sig"`
	Start              time.Time           `json:"start"`
	StartTzId          string              `json:"startTzId"`
	Summary            string              `json:"summary"`
	Tags               string              `gorm:"type:jsonb" json:"tags"`
	Title              string              `json:"title"`
	Type               CalendarEventType   `json:"type"`
	URL                string              `json:"url"`
}

// Check if CalendarEventType matches a string
func (c CalendarEventType) Matches(s string) bool {
	return string(c) == s
}

// Return string for CalendarEventType
func (c CalendarEventType) String() string {
	return string(c)
}
