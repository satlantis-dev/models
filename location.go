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
	ID                    uint              `gorm:"primaryKey" json:"id"`
	Name                  string            `json:"name"`
	CreatedAt             time.Time         `json:"-"`
	UpdatedAt             time.Time         `json:"-"`
	DeletedAt             *time.Time        `gorm:"index" json:"-,omitempty"`
	Accounts              []LocationAccount `gorm:"foreignKey:LocationID" json:"accounts"`
	Lat                   float64           `json:"lat"`
	Lng                   float64           `json:"lng"`
	LocationTags          []LocationTag     `gorm:"many2many:location_location_tags" json:"locationTags"`
	OSMRef                string            `gorm:"uniqueIndex;not null" json:"osmRef"`
	GoogleID              string            `gorm:"uniqueIndex;not null" json:"googleId"`
	PlaceID               uint              `gorm:"index" json:"placeId"`
	EventID               *uint             `gorm:"index" json:"eventId"`
	Event                 Event             `json:"event"`
	Notes                 []LocationNote    `gorm:"foreignKey:LocationID" json:"notes"`
	Score                 float64           `json:"score"`
	Address               Address           `gorm:"type:jsonb" json:"address"`
	Phone                 string            `json:"phone"`
	BusinessStatus        BusinessStatus    `gorm:"type:text" json:"businessStatus"`
	OpeningHours          OpeningHours      `gorm:"type:jsonb" json:"openingHours"`
	PriceLevel            PriceLevel        `json:"priceLevel"`
	GoogleRating          float64           `json:"googleRating"`
	GoogleUserRatingCount int               `json:"googleUserRatingCount"`
	GoogleMapsUrl         string            `json:"googleMapsUrl"`
	WebsiteUrl            string            `json:"websiteUrl"`
	IsClaimed             bool              `json:"isClaimed"`
	Bio                   *string           `json:"bio"`
}

// LocationDTO
type LocationDTO struct {
	ID             uint          `json:"id"`
	Name           string        `json:"name"`
	Accounts       []AccountDTO  `json:"accounts"`
	Lat            float64       `json:"lat"`
	Lng            float64       `json:"lng"`
	LocationTags   []LocationTag `gorm:"many2many:location_location_tags" json:"locationTags"`
	PlaceID        uint          `json:"placeId"`
	PublicLocation bool          `json:"publicLocation"`
	Score          float64       `json:"score"`
}

func (l Location) ToDTO(db *gorm.DB) (*LocationDTO, error) {
	var accounts []Account
	// Get the accounts
	err := db.Table("location_accounts").Select("accounts.*").
		Joins("join accounts on accounts.id = location_accounts.account_id").
		Where("location_accounts.location_id = ?", l.ID).
		Scan(&accounts).Error
	if err != nil {
		return nil, err
	}

	accountDTOs := make([]AccountDTO, len(accounts))

	for i, account := range accounts {
		accountDTOs[i] = account.ToDTO()
	}

	return &LocationDTO{
		ID:           l.ID,
		Accounts:     accountDTOs,
		Lat:          l.Lat,
		Lng:          l.Lng,
		LocationTags: l.LocationTags,
		PlaceID:      l.PlaceID,
		Score:        l.Score,
	}, nil
}
