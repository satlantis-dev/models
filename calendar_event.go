package models

import "time"

type CalendarEventType int

const (
	Conference CalendarEventType = iota + 1
	Meetup
	Hackathon
	Concert
	Workshop
	Party
	Play
	Sports
	Exhibition
	Festival
	Music
	Other
)

func (c CalendarEventType) Hashtag() string {
	names := [...]string{
		"#conference",
		"#meetup",
		"#hackathon",
		"#concert",
		"#workshop",
		"#party",
		"#play",
		"#sports",
		"#exhibition",
		"#festival",
		"#music",
		"#other",
	}

	if c < Conference || c > Other {
		return "Unknown"
	}

	return names[c-1]
}

type CalendarEvent struct {
	ID                 uint                `gorm:"primaryKey" json:"id"`
	AccountID          uint                `gorm:"index" json:"accountId"`
	Account            Account             `json:"account"`
	Announcements      []CalendarEventNote `gorm:"foreignKey:CalendarEventID" json:"announcements"`
	CreatedAt          *time.Time          `json:"createdAt"`
	CalendarEventRSVPs []CalendarEventRSVP `gorm:"foreignKey:CalendarEventID" json:"calendarEventRsvps"`
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
