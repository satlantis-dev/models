package database

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

// Not a GORM model because it's not stored in the database
type CalendarEvent struct {
	ID            string
	Description   string
	End           time.Time
	Geohash       string
	Image         string
	Location      string
	Start         time.Time
	StartTimezone string
	Title         string
	Type          CalendarEventType
	URL           string
}
