package models

import (
	"time"

	"github.com/lib/pq"
)

type InterestCategory int

const (
	GeneralInterest InterestCategory = iota + 1
	LocationInterest
	ActivityInterest
	FoodInterest
)

type Interest struct {
	ID                    uint             `gorm:"primaryKey" json:"id"`
	Name                  string           `gorm:"unique" json:"name"`
	CreatedAt             time.Time        `json:"-"`
	UpdatedAt             time.Time        `json:"-"`
	DeletedAt             *time.Time       `gorm:"index" json:"-,omitempty"`
	Description           string           `gorm:"type:text" json:"description"`
	RecommendationsByNpub pq.StringArray   `gorm:"type:varchar[]" json:"recommendationsByNpub"`
	RecommendationsById   pq.Int32Array    `gorm:"type:integer[]" json:"recommendationsById"`
	AutofollowsByNpub     pq.StringArray   `gorm:"type:varchar[]" json:"autofollowsByNpub"`
	AutofollowsById       pq.Int32Array    `gorm:"type:integer[]" json:"autofollowsById"`
	Hashtags              pq.StringArray   `gorm:"type:varchar[]" json:"hashtags"`
	LocationTags          []LocationTag    `gorm:"many2many:interest_location_tags;constraint:OnDelete:CASCADE;" json:"locationTags"`
	Category              InterestCategory `json:"category"`
	Section               string           `json:"section"`
	Emoji                 string           `gorm:"type:varchar(7)" json:"emoji"`
	ContentUse            bool             `json:"contentUse"`
	LocationUse           bool             `json:"locationUse"`
	EventUse              bool             `json:"eventUse"`
	PeopleUse             bool             `json:"peopleUse"`
}

type WeightedInterest struct {
	InterestID uint
	Weight     float64
}
