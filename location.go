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
	b, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return string(b), nil
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
	b, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return string(b), nil
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
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

type ReviewHighlights struct {
	Amenities     *string `json:"amenities,omitempty"`
	Cleanliness   *string `json:"cleanliness,omitempty"`
	Experience    *string `json:"experience,omitempty"`
	Facilities    *string `json:"facilities,omitempty"`
	FoodAndDrinks *string `json:"food_and_drinks,omitempty"`
	Location      *string `json:"location,omitempty"`
	Music         *string `json:"music,omitempty"`
	Rooms         *string `json:"rooms,omitempty"`
	Service       *string `json:"service,omitempty"`
	ValueForMoney *string `json:"value_for_money,omitempty"`
	Vibe          *string `json:"vibe,omitempty"`
}

func (r *ReviewHighlights) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, r)
}

func (r ReviewHighlights) Value() (driver.Value, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

type Socials struct {
	Facebook  *string `json:"facebook,omitempty"`  // username -> "https://facebook.com/<username>"
	Instagram *string `json:"instagram,omitempty"` // username -> "https://instagram.com/<username>"
	Line      *string `json:"line,omitempty"`      // business id -> "https://page.line.me/<business id>"
	Telegram  *string `json:"telegram,omitempty"`  // username -> "https://t.me/<username>"
	TikTok    *string `json:"tiktok,omitempty"`    // username -> "https://www.tiktok.com/@<username>"
	WhatsApp  *string `json:"whatsapp,omitempty"`  // intl phone number -> "https://wa.me/<intl phone number>"
	X         *string `json:"x,omitempty"`         // username -> "https://x.com/<username>"
	YouTube   *string `json:"youtube,omitempty"`   // username -> "https://youtube.com/@<username>"
}

func (s *Socials) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, s)
}

func (s Socials) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

type Location struct {
	ID                    uint                   `gorm:"primaryKey" json:"id"`
	CreatedAt             time.Time              `json:"-"`
	UpdatedAt             time.Time              `json:"-"`
	DeletedAt             *time.Time             `gorm:"index" json:"-,omitempty"`
	AccountRoles          []AccountLocationRole  `gorm:"foreignKey:LocationID" json:"accountRoles"`
	Address               Address                `gorm:"type:jsonb;serializer:json" json:"address"`
	Bio                   *string                `json:"bio"`
	BusinessStatus        BusinessStatus         `gorm:"type:text" json:"businessStatus"`
	Claim                 LocationClaim          `gorm:"foreignKey:LocationID" json:"claim"`
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
	OpeningHours          OpeningHours           `gorm:"type:jsonb;serializer:json" json:"openingHours"`
	OSMRef                string                 `json:"osmRef"`
	Phone                 string                 `json:"phone"`
	PriceLevel            PriceLevel             `json:"priceLevel"`
	Score                 float64                `json:"score"`
	TripadvisorRating     ExternalRating         `gorm:"type:jsonb;serializer:json" json:"tripadvisorRating"`
	GooglePlacesRating    ExternalRating         `gorm:"type:jsonb;serializer:json" json:"googlePlacesRating"`
	WebsiteUrl            string                 `json:"websiteUrl"`
	Email                 string                 `json:"email"`
	ReviewSummary         string                 `json:"reviewSummary"`
	ReviewHighlights      ReviewHighlights       `gorm:"type:jsonb;serializer:json" json:"reviewHighlights"`
	Socials               Socials                `gorm:"type:jsonb;serializer:json" json:"socials"`
}

// LocationDTO
type LocationDTO struct {
	ID              uint          `json:"id"`
	GoogleID        string        `gorm:"uniqueIndex;not null" json:"googleId"`
	Address         Address       `gorm:"type:jsonb;serializer:json" json:"address"`
	Bio             *string       `json:"bio"`
	Email           string        `json:"email"`
	Rating          float64       `json:"rating"`
	UserRatingCount int           `json:"userRatingCount"`
	GoogleMapsUrl   string        `json:"googleMapsUrl"`
	Hook            *string       `json:"hook"`
	Image           string        `json:"image"`
	IsClaimed       bool          `json:"isClaimed"`
	Lat             float64       `json:"lat"`
	Lng             float64       `json:"lng"`
	LocationTags    []LocationTag `gorm:"many2many:location_location_tags" json:"locationTags"`
	Name            string        `json:"name"`
	OpeningHours    OpeningHours  `gorm:"type:jsonb;serializer:json" json:"openingHours"`
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
		GoogleID:        l.GoogleID,
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
		OpeningHours:    l.OpeningHours,
		PlaceID:         l.PlaceID,
		PlaceOSMRef:     l.Place.OSMRef,
		ReviewSummary:   l.ReviewSummary,
	}, nil
}

type NearbyLocationDTO struct {
	LocationDTO
	Distance int `json:"distance"`
}
