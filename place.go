package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
)

type PlaceLevel string

const (
	PlaceLevelRegion       PlaceLevel = "region"
	PlaceLevelCity         PlaceLevel = "city"
	PlaceLevelNeighborhood PlaceLevel = "neighborhood"
)

type BoundingBox struct {
	MinLat float64 `json:"minlat"`
	MaxLat float64 `json:"maxlat"`
	MinLng float64 `json:"minlng"`
	MaxLng float64 `json:"maxlng"`
}

func (b *BoundingBox) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &b)
}

func (b BoundingBox) Value() (driver.Value, error) {
	return json.Marshal(b)
}

type Place struct {
	ID             uint                 `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time            `json:"-"`
	UpdatedAt      time.Time            `json:"-"`
	DeletedAt      *time.Time           `gorm:"index" json:"-,omitempty"`
	AccountRoles   []AccountPlaceRole   `gorm:"foreignKey:PlaceID" json:"accountRoles"`
	Active         bool                 `json:"active"`
	Banner         string               `gorm:"type:text" json:"banner"`
	BoundingBox    BoundingBox          `gorm:"type:jsonb" json:"boundingBox"`
	CategoryScores []PlaceCategoryScore `gorm:"foreignKey:PlaceID" json:"categoryScores"`
	CountryID      uint                 `gorm:"index" json:"countryId"`
	Country        Country              `json:"country"`
	Descendants    []PlaceWithClosure   `gorm:"-" json:"descendants"`
	Description    string               `gorm:"type:text" json:"description"`
	EventID        *uint                `gorm:"index" json:"eventId"`
	Event          Event                `json:"event"`
	Lat            float64              `json:"lat"`
	Level          PlaceLevel           `gorm:"type:text" json:"level"`
	Lng            float64              `json:"lng"`
	Metrics        []PlaceMetric        `gorm:"foreignKey:PlaceID" json:"metrics"`
	Name           string               `gorm:"index;type:text" json:"name"`
	Notes          []PlaceNote          `gorm:"foreignKey:PlaceID" json:"notes"`
	OSMID          *uint                `json:"osmId"`
	OSMLevel       string               `json:"osmLevel"`
	OSMType        OSMType              `json:"osmType"`
	OSMRef         string               `gorm:"uniqueIndex" json:"osmRef"`
	RegionID       *uint                `gorm:"index" json:"regionId"`
	Region         Region               `gorm:"foreignKey:RegionID" json:"region"`
	Slug           string               `gorm:"type:text" json:"slug"` // Unique slug for the place navigation
	WeatherID      *uint                `gorm:"index" json:"weatherId"`
	Weather        Weather              `gorm:"foreignKey:PlaceID" json:"weather"`
	Hashtags       pq.StringArray       `gorm:"type:varchar[]" json:"hashtags"`
}

// Place With Closure
type PlaceWithClosure struct {
	Place
	AncestorID   uint `gorm:"index" json:"ancestorId"`
	DescendantID uint `gorm:"index" json:"descendantId"`
	Depth        uint `json:"depth"`
}

type PlaceDTO struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Active      bool           `json:"active"`
	Banner      string         `gorm:"type:text" json:"banner"`
	CountryID   uint           `gorm:"index" json:"countryId"`
	Country     Country        `json:"country"`
	Description string         `gorm:"type:text" json:"description"`
	EventID     *uint          `gorm:"index" json:"eventId"`
	Lat         float64        `json:"lat"`
	Level       PlaceLevel     `gorm:"type:text" json:"level"`
	Lng         float64        `json:"lng"`
	Name        string         `gorm:"index;type:text" json:"name"`
	OSMID       *uint          `json:"osmId"`
	OSMLevel    string         `json:"osmLevel"`
	OSMType     OSMType        `json:"osmType"`
	OSMRef      string         `gorm:"uniqueIndex" json:"osmRef"`
	RegionID    *uint          `gorm:"index" json:"regionId"`
	Slug        string         `gorm:"type:text" json:"slug"` // Unique slug for the place navigation
	WeatherID   *uint          `gorm:"index" json:"weatherId"`
	Weather     Weather        `json:"weather"`
	Hashtags    pq.StringArray `gorm:"type:varchar[]" json:"hashtags"`
}

// ToDTO - Convert Place to PlaceDTO.
func (place Place) ToDTO() *PlaceDTO {
	return &PlaceDTO{
		ID:          place.ID,
		Active:      place.Active,
		Banner:      place.Banner,
		CountryID:   place.CountryID,
		Country:     place.Country,
		Description: place.Description,
		EventID:     place.EventID,
		Lat:         place.Lat,
		Level:       place.Level,
		Lng:         place.Lng,
		Name:        place.Name,
		OSMID:       place.OSMID,
		OSMLevel:    place.OSMLevel,
		OSMType:     place.OSMType,
		OSMRef:      place.OSMRef,
		RegionID:    place.RegionID,
		Slug:        place.Slug,
		WeatherID:   place.WeatherID,
		Weather:     place.Weather,
		Hashtags:    place.Hashtags,
	}
}
