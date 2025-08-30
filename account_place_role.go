package models

import (
	"time"

	"gorm.io/gorm"
)

type AccountPlaceRoleType int

const (
	Follower AccountPlaceRoleType = iota + 1
	Visitor
	Inhabitant
	Ambassador
)

type AccountPlaceRole struct {
	AccountID         uint                 `gorm:"index;uniqueIndex:idx_account_place" json:"accountId"`
	Account           *Account             `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE" json:"account,omitempty"`
	PlaceID           uint                 `gorm:"index;uniqueIndex:idx_account_place" json:"placeId"`
	Place             *Place               `gorm:"foreignKey:PlaceID;constraint:OnDelete:CASCADE" json:"place,omitempty"`
	Type              AccountPlaceRoleType `gorm:"not null" json:"type"`
	AmbassadorRequest bool                 `json:"ambassadorRequest"`
	CreatedAt         time.Time            `json:"-"`
	UpdatedAt         time.Time            `json:"-"`
	DeletedAt         gorm.DeletedAt       `gorm:"index" json:"-"`
}
