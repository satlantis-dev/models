package models

import "time"

type CalendarEvent struct {
	ID                 uint                        `gorm:"primaryKey" json:"id"`
	AccountID          uint                        `gorm:"index" json:"accountId"`
	Account            Account                     `json:"account"`
	Announcements      []CalendarEventAccouncement `gorm:"foreignKey:CalendarEventID" json:"announcements"`
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
	Type               string                      `json:"type"`
	URL                string                      `json:"url"`
}
