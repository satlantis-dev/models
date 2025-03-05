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
	Formatted    string `json:"formatted"`
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

type ExternalRating struct {
	Source string  `json:"source"`
	Id     string  `json:"id"`
	Rating float64 `json:"rating"`
	Count  int     `json:"count"`
	Link   string  `json:"link"`
}

func (r *ExternalRating) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, r)
}

func (r ExternalRating) Value() (driver.Value, error) {
	return json.Marshal(r)
}

type Location struct {
	ID                    uint                   `gorm:"primaryKey" json:"id"`
	CreatedAt             time.Time              `json:"-"`
	UpdatedAt             time.Time              `json:"-"`
	DeletedAt             *time.Time             `gorm:"index" json:"-,omitempty"`
	AccountRoles          []AccountLocationRole  `gorm:"foreignKey:LocationID" json:"accountRoles"`
	Address               Address                `gorm:"type:jsonb" json:"address"`
	Bio                   *string                `json:"bio"`
	BusinessStatus        BusinessStatus         `gorm:"type:text" json:"businessStatus"`
	Claim                 LocationClaim          `gorm:"foreignKey:LocationID" json:"claim"`
	EventID               *uint                  `gorm:"index" json:"eventId"`
	Event                 Event                  `json:"event"`
	GoogleID              string                 `gorm:"uniqueIndex;not null" json:"googleId"`
	GoogleMapsUrl         string                 `json:"googleMapsUrl"`
	Rating                float64                `json:"rating"`
	UserRatingCount       int                    `json:"userRatingCount"`
	Hook                  *string                `json:"hook"`
	Image                 string                 `json:"image"`
	IsClaimed             bool                   `json:"isClaimed"`
	Lat                   float64                `json:"lat"`
	Lng                   float64                `json:"lng"`
	LocationGalleryImages []LocationGalleryImage `gorm:"foreignKey:LocationID" json:"locationGalleryImages,omitempty"`
	LocationTags          []LocationTag          `gorm:"many2many:location_location_tags" json:"locationTags"`
	PlaceID               uint                   `gorm:"index" json:"placeId"`
	Place                 Place                  `json:"place"`
	Name                  string                 `json:"name"`
	Notes                 []LocationNote         `gorm:"foreignKey:LocationID" json:"notes"`
	OpeningHours          OpeningHours           `gorm:"type:jsonb" json:"openingHours"`
	OSMRef                string                 `gorm:"uniqueIndex;not null" json:"osmRef"`
	Phone                 string                 `json:"phone"`
	PriceLevel            PriceLevel             `json:"priceLevel"`
	Score                 float64                `json:"score"`
	TripadvisorRating     ExternalRating         `gorm:"type:jsonb" json:"tripadvisorRating"`
	GooglePlacesRating    ExternalRating         `gorm:"type:jsonb" json:"googlePlacesRating"`
	WebsiteUrl            string                 `json:"websiteUrl"`
	Email                 string                 `json:"email"`
	ReviewSummary         string                 `json:"reviewSummary"`
}

// LocationDTO
type LocationDTO struct {
	ID              uint          `json:"id"`
	Address         Address       `json:"address"`
	Bio             *string       `json:"bio"`
	Email           string        `json:"email"`
	Rating          float64       `json:"Rating"`
	UserRatingCount int           `json:"UserRatingCount"`
	GoogleMapsUrl   string        `json:"googleMapsUrl"`
	Hook            *string       `gorm:"size:70" json:"hook"`
	Image           string        `json:"image"`
	IsClaimed       bool          `json:"isClaimed"`
	Lat             float64       `json:"lat"`
	Lng             float64       `json:"lng"`
	LocationTags    []LocationTag `gorm:"many2many:location_location_tags" json:"locationTags"`
	Name            string        `json:"name"`
	OpeningHours    OpeningHours  `json:"openingHours"`
	OSMRef          string        `json:"osmRef"`
	PlaceID         uint          `json:"placeId"`
	PlaceOSMRef     string        `json:"placeOsmRef"`
	ReviewSummary   string        `json:"reviewSummary"`
}

func (LocationDTO) TableName() string {
	return "locations"
}

func (l Location) ToDTO(db *gorm.DB) (*LocationDTO, error) {

	return &LocationDTO{
		ID:              l.ID,
		Address:         l.Address,
		Bio:             l.Bio,
		Email:           l.Email,
		Rating:          l.Rating,
		UserRatingCount: l.UserRatingCount,
		GoogleMapsUrl:   l.GoogleMapsUrl,
		Hook:            l.Hook,
		Image:           l.Image,
		IsClaimed:       l.IsClaimed,
		Lat:             l.Lat,
		Lng:             l.Lng,
		LocationTags:    l.LocationTags,
		Name:            l.Name,
		OSMRef:          l.OSMRef,
		OpeningHours:    l.OpeningHours,
		PlaceID:         l.PlaceID,
		PlaceOSMRef:     l.Place.OSMRef,
		ReviewSummary:   l.ReviewSummary,
	}, nil
}
