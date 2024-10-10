package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BusinessStatus string

const (
	BusinessStatusClosedPerm  BusinessStatus = "CLOSED_PERMANENTLY"
	BusinessStatusClosedTemp  BusinessStatus = "CLOSED_TEMPORARILY"
	BusinessStatusOperational BusinessStatus = "OPERATIONAL"
)

type PriceLevel string

const (
	PriceLevelX PriceLevel = "PRICE_LEVEL_UNSPECIFIED"
	PriceLevel0 PriceLevel = "PRICE_LEVEL_FREE"
	PriceLevel1 PriceLevel = "PRICE_LEVEL_INEXPENSIVE"
	PriceLevel2 PriceLevel = "PRICE_LEVEL_MODERATE"
	PriceLevel3 PriceLevel = "PRICE_LEVEL_EXPENSIVE"
	PriceLevel4 PriceLevel = "PRICE_LEVEL_VERY_EXPENSIVE"
)

type Address struct {
	StreetNumber string `json:"streetNumber"`
	Route        string `json:"route"`
	Locality     string `json:"locality"`
	PostalCode   string `json:"postalCode"`
	Country      string `json:"country"`
}

func (a *Address) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}

	return json.Unmarshal(b, a)
}
func (a Address) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type OpeningHours struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

func (o *OpeningHours) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, o)
}

func (o OpeningHours) Value() (driver.Value, error) {
	return json.Marshal(o)
}

type Location struct {
	ID                    uint           `gorm:"primaryKey" json:"id"`
	CreatedAt             time.Time      `json:"-"`
	UpdatedAt             time.Time      `json:"-"`
	DeletedAt             *time.Time     `gorm:"index" json:"-,omitempty"`
	Address               Address        `gorm:"type:jsonb" json:"address"`
	Bio                   *string        `json:"bio"`
	BusinessStatus        BusinessStatus `gorm:"type:text" json:"businessStatus"`
	EventID               *uint          `gorm:"index" json:"eventId"`
	Event                 Event          `json:"event"`
	GoogleID              string         `gorm:"uniqueIndex;not null" json:"googleId"`
	GoogleMapsUrl         string         `json:"googleMapsUrl"`
	GoogleRating          float64        `json:"googleRating"`
	GoogleUserRatingCount int            `json:"googleUserRatingCount"`
	Image                 string         `json:"image"`
	IsClaimed             bool           `json:"isClaimed"`
	Lat                   float64        `json:"lat"`
	Lng                   float64        `json:"lng"`
	LocationTags          []LocationTag  `gorm:"many2many:location_location_tags" json:"locationTags"`
	PlaceID               uint           `gorm:"index" json:"placeId"`
	Place                 Place          `json:"place"`
	Name                  string         `json:"name"`
	Notes                 []LocationNote `gorm:"foreignKey:LocationID" json:"notes"`
	OpeningHours          OpeningHours   `gorm:"type:jsonb" json:"openingHours"`
	OSMRef                string         `gorm:"uniqueIndex;not null" json:"osmRef"`
	Phone                 string         `json:"phone"`
	PriceLevel            PriceLevel     `json:"priceLevel"`
	Score                 float64        `json:"score"`
	WebsiteUrl            string         `json:"websiteUrl"`
	Email                 string         `json:"email"`
}

// LocationDTO
type LocationDTO struct {
	ID                    uint          `json:"id"`
	Bio                   *string       `json:"bio"`
	Image                 string        `json:"image"`
	Lat                   float64       `json:"lat"`
	Lng                   float64       `json:"lng"`
	LocationTags          []LocationTag `gorm:"many2many:location_location_tags" json:"locationTags"`
	Name                  string        `json:"name"`
	Email                 string        `json:"email"`
	PlaceID               uint          `json:"placeId"`
	Score                 float64       `json:"score"`
	OSMRef                string        `json:"osmRef"`
	PlaceOSMRef           string        `json:"placeOsmRef"`
	OpeningHours          OpeningHours  `json:"openingHours"`
	Address               Address       `json:"address"`
	GoogleRating          float64       `json:"googleRating"`
	GoogleUserRatingCount int           `json:"googleUserRatingCount"`
}

func (l Location) ToDTO(db *gorm.DB) (*LocationDTO, error) {
	// Get the place
	var place Place
	_ = db.First(&place, l.PlaceID).Select("OSMRef").Error

	return &LocationDTO{
		ID:                    l.ID,
		Bio:                   l.Bio,
		Image:                 l.Image,
		Lat:                   l.Lat,
		Lng:                   l.Lng,
		LocationTags:          l.LocationTags,
		Name:                  l.Name,
		Email:                 l.Email,
		PlaceID:               l.PlaceID,
		Score:                 l.Score,
		OSMRef:                l.OSMRef,
		PlaceOSMRef:           place.OSMRef,
		OpeningHours:          l.OpeningHours,
		Address:               l.Address,
		GoogleRating:          l.GoogleRating,
		GoogleUserRatingCount: l.GoogleUserRatingCount,
	}, nil
}
