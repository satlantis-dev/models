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
	ID                 uint                  `gorm:"primaryKey" json:"id"`
	CreatedAt          time.Time             `json:"-"`
	UpdatedAt          time.Time             `json:"-"`
	DeletedAt          *time.Time            `gorm:"index" json:"-,omitempty"`
	AccountRoles       *[]AccountPlaceRole   `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE;" json:"accountRoles,omitempty"`
	Active             bool                  `json:"active"`
	Banner             string                `gorm:"type:text" json:"banner"`
	BoundingBox        BoundingBox           `gorm:"type:jsonb" json:"boundingBox"`
	CategoryScores     *[]PlaceCategoryScore `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE;" json:"categoryScores,omitempty"`
	CountryID          uint                  `gorm:"index" json:"countryId"`
	Country            Country               `json:"country"`
	Descendants        []PlaceWithClosure    `gorm:"-" json:"descendants"`
	Description        string                `gorm:"type:text" json:"description"`
	EventID            *uint                 `gorm:"index" json:"eventId"`
	Event              *Event                `json:"event,omitempty"`
	Featured           bool                  `json:"featured"`
	Lat                float64               `json:"lat"`
	Level              PlaceLevel            `gorm:"type:text" json:"level"`
	Lng                float64               `json:"lng"`
	Metrics            *[]PlaceMetric        `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE;" json:"metrics,omitempty"`
	Name               string                `gorm:"index;type:text" json:"name"`
	Notes              *[]PlaceNote          `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE;" json:"notes,omitempty"`
	OSMID              *uint                 `json:"osmId"`
	OSMLevel           string                `json:"osmLevel"`
	OSMType            OSMType               `json:"osmType"`
	OSMRef             string                `gorm:"uniqueIndex" json:"osmRef"`
	PlaceGalleryImages *[]PlaceGalleryImage  `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE;" json:"placeGalleryImages,omitempty"`
	RegionID           *uint                 `gorm:"index" json:"regionId"`
	Region             *Region               `gorm:"foreignKey:RegionID" json:"region,omitempty"`
	Timezone           string                `gorm:"type:text" json:"timezone"`
	WeatherID          *uint                 `gorm:"index" json:"weatherId"`
	Weather            *Weather              `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE;" json:"weather,omitempty"`
	Hashtags           pq.StringArray        `gorm:"type:varchar[]" json:"hashtags"`
}

// Place With Closure
type PlaceWithClosure struct {
	Place
	AncestorID   uint `gorm:"index" json:"ancestorId"`
	DescendantID uint `gorm:"index" json:"descendantId"`
	Depth        uint `json:"depth"`
}

type PlaceDTO struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Banner    string     `gorm:"type:text" json:"banner"`
	CountryID uint       `gorm:"index" json:"countryId"`
	Country   Country    `json:"country"`
	Level     PlaceLevel `gorm:"type:text" json:"level"`
	Name      string     `gorm:"index;type:text" json:"name"`
	OSMID     *uint      `json:"osmId"`
	OSMLevel  string     `json:"osmLevel"`
	OSMType   OSMType    `json:"osmType"`
	OSMRef    string     `gorm:"uniqueIndex" json:"osmRef"`
}

func (PlaceDTO) TableName() string {
	return "places"
}

// ToDTO - Convert Place to PlaceDTO.
func (place Place) ToDTO() *PlaceDTO {
	return &PlaceDTO{
		ID:        place.ID,
		Banner:    place.Banner,
		CountryID: place.CountryID,
		Country:   place.Country,
		Level:     place.Level,
		Name:      place.Name,
		OSMID:     place.OSMID,
		OSMLevel:  place.OSMLevel,
		OSMType:   place.OSMType,
		OSMRef:    place.OSMRef,
	}
}
