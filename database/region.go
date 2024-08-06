package database

import (
	"time"

	"github.com/lib/pq"
)

type Region struct {
	ID             uint                  `gorm:"primaryKey" json:"id"`
	Code           string                `gorm:"type:text" json:"code"`
	CreatedAt      time.Time             `json:"-"`
	UpdatedAt      time.Time             `json:"-"`
	DeletedAt      *time.Time            `gorm:"index" json:"-,omitempty"`
	Banner         string                `gorm:"type:text" json:"banner"`
	CategoryScores []RegionCategoryScore `gorm:"foreignKey:RegionID" json:"categoryScores"`
	CountryID      uint                  `gorm:"index" json:"countryId"`
	Country        Country               `json:"country"`
	Description    string                `gorm:"type:text" json:"description"`
	EventID        *uint                 `gorm:"index" json:"eventId"`
	Event          Event                 `json:"event"`
	Lat            float64               `json:"lat"`
	Lng            float64               `json:"lng"`
	Metrics        []RegionMetric        `gorm:"foreignKey:RegionID" json:"metrics"`
	Name           string                `gorm:"index;type:text" json:"name"`
	OSMID          *uint                 `json:"osmId"`
	OSMLevel       string                `json:"osmLevel"`
	OSMType        OSMType               `json:"osmType"`
	OSMRef         string                `gorm:"uniqueIndex" json:"osmRef"`
	Places         []Place               `gorm:"foreignKey:RegionID" json:"places"`
	Slug           string                `gorm:"type:text" json:"slug"` // Unique slug for the region navigation
	Hashtags       pq.StringArray        `gorm:"type:varchar[]" json:"hashtags"`
}
