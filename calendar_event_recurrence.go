package models

import (
	"time"

	"gorm.io/gorm"
)

// CalendarEventSeries is the recurring rule + ownership context.
type CalendarEventSeries struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	AccountID uint           `gorm:"not null;index" json:"accountId"`
	Account   *AccountDTO    `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Anchor event holds base event data (title, summary, time, location, etc).
	AnchorEventID uint           `gorm:"index" json:"anchorEventId"`
	AnchorEvent   *CalendarEvent `gorm:"foreignKey:AnchorEventID;constraint:OnDelete:SET NULL" json:"anchorEvent,omitempty"`

	// RFC5545 RRULE (e.g. FREQ=WEEKLY;BYDAY=MO,WE;INTERVAL=1) and GeneratedThrough timestamp for tracking generation progress.
	RRule            string     `gorm:"type:text;not null" json:"rrule"`
	GeneratedThrough *time.Time `gorm:"index" json:"generatedThrough,omitempty"`
}

func (CalendarEventSeries) TableName() string {
	return "calendar_event_series"
}

// CalendarEventOccurrence is optional cache/materialization for fast range queries.
type CalendarEventOccurrence struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Link back to series
	SeriesID uint                 `gorm:"not null;index;uniqueIndex:idx_calendar_event_occurrence_series_original_start" json:"seriesId"`
	Series   *CalendarEventSeries `gorm:"foreignKey:SeriesID;constraint:OnDelete:CASCADE" json:"series,omitempty"`

	// Original start time in UTC for this occurrence (used for uniqueness e.g. when generating occurrences).
	OriginalStartUTC time.Time `gorm:"not null;index;uniqueIndex:idx_calendar_event_occurrence_series_original_start" json:"originalStartUtc"`

	// Effective event row for rendering (anchor event by default, event override if modified).
	CalendarEventID uint           `gorm:"index" json:"calendarEventId"`
	CalendarEvent   *CalendarEvent `gorm:"foreignKey:CalendarEventID;constraint:OnDelete:SET NULL" json:"calendarEvent,omitempty"`

	// Recurrence instance time data (for fast querying/rendering).
	Start     time.Time `gorm:"not null;index" json:"start"`
	StartTzId string    `gorm:"not null" json:"startTzId"`
	End       time.Time `gorm:"not null;index" json:"end"`
	EndTzId   string    `gorm:"not null" json:"endTzId"`

	// Cancellation flag (if occurrence is cancelled)
	IsCancelled bool `gorm:"not null;default:false;index" json:"isCancelled"`
}

func (CalendarEventOccurrence) TableName() string {
	return "calendar_event_occurrences"
}
