package models

import (
	"time"

	"gorm.io/gorm"
)

type Calendar struct {
	ID                 uint                     `json:"id"`
	AccountID          uint                     `json:"account_id"`
	Account            *Account                 `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	Name               string                   `json:"name"`
	Description        string                   `json:"description"`
	Slug               string                   `gorm:"size:70" json:"slug"`
	Banner             string                   `json:"banner"`
	CreatedAt          time.Time                `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time                `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt          *gorm.DeletedAt          `gorm:"index" json:"-"`
	CalendarEvents     *[]CalendarCalendarEvent `gorm:"foreignKey:CalendarID;constraint:OnDelete:CASCADE;" json:"calendarEvents,omitempty"`
	Events             *[]CalendarEvent         `gorm:"many2many:calendar_calendar_events;constraint:OnDelete:CASCADE;" json:"events,omitempty"`
	EventCount         int                      `gorm:"-" json:"eventCount"`
	UpcomingEventCount int                      `gorm:"-" json:"upcomingEventCount"`
	EventTags          *[]CalendarEventTag      `gorm:"many2many:calendar_calendar_event_tags;constraint:OnDelete:CASCADE" json:"eventTags,omitempty"`
	Contributors       *[]AccountMiniDTO        `gorm:"-" json:"contributors,omitempty"`
	NumSubscriptions   int                      `gorm:"-" json:"numSubscriptions"`
	PlaceID            *uint                    `json:"placeId,omitempty"`
	Place              *Place                   `gorm:"foreignKey:PlaceID;constraint:OnDelete:SET NULL" json:"place,omitempty"`
	IsPublic           bool                     `gorm:"default:true" json:"isPublic"`
	Featured           bool                     `gorm:"default:false;index:idx_calendar_featured,where:featured = true" json:"featured"`
	CommunityID        *uint                    `gorm:"index" json:"communityId,omitempty"`
	Community          *Community               `gorm:"foreignKey:CommunityID;constraint:OnDelete:SET NULL" json:"community,omitempty"`
}

// CalendarDTO is the trimmed calendar shape the list endpoints return (discover, an account's
// calendars and subscriptions). It keeps the keys of Calendar that lists fill and adds the owner as
// a mini account, instead of the full Account relation. Single-calendar endpoints still return
// Calendar.
type CalendarDTO struct {
	ID                 uint                `json:"id"`
	AccountID          uint                `json:"account_id"`
	Owner              *AccountMiniDTO     `json:"owner,omitempty"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	Slug               string              `json:"slug"`
	Banner             string              `json:"banner"`
	CreatedAt          time.Time           `json:"createdAt"`
	EventCount         int                 `json:"eventCount"`
	UpcomingEventCount int                 `json:"upcomingEventCount"`
	EventTags          *[]CalendarEventTag `json:"eventTags,omitempty"`
	NumSubscriptions   int                 `json:"numSubscriptions"`
	PlaceID            *uint               `json:"placeId,omitempty"`
	IsPublic           bool                `json:"isPublic"`
	Featured           bool                `json:"featured"`
	CommunityID        *uint               `json:"communityId,omitempty"`
}

func (CalendarDTO) TableName() string {
	return "calendars"
}

// ToDTO converts the calendar to its list shape. The owner comes from Account when it is loaded.
func (c *Calendar) ToDTO() CalendarDTO {
	dto := CalendarDTO{
		ID:                 c.ID,
		AccountID:          c.AccountID,
		Name:               c.Name,
		Description:        c.Description,
		Slug:               c.Slug,
		Banner:             c.Banner,
		CreatedAt:          c.CreatedAt,
		EventCount:         c.EventCount,
		UpcomingEventCount: c.UpcomingEventCount,
		EventTags:          c.EventTags,
		NumSubscriptions:   c.NumSubscriptions,
		PlaceID:            c.PlaceID,
		IsPublic:           c.IsPublic,
		Featured:           c.Featured,
		CommunityID:        c.CommunityID,
	}
	if c.Account != nil {
		owner := c.Account.ToMiniDTO()
		dto.Owner = &owner
	}
	return dto
}

// ToCalendarDTOs converts a slice of calendars to their list shape.
func ToCalendarDTOs(calendars []Calendar) []CalendarDTO {
	dtos := make([]CalendarDTO, len(calendars))
	for i := range calendars {
		dtos[i] = calendars[i].ToDTO()
	}
	return dtos
}
