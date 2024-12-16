package models

import (
	"time"
)

type AccountPlaceRoleType int

const (
	Follower AccountPlaceRoleType = iota + 1
	Visitor
	Inhabitant
	Ambassador
)

type AccountPlaceRole struct {
	AccountID         uint                 `gorm:"index;primaryKey" json:"accountId"`
	Account           AccountDTO           `gorm:"foreignKey:AccountID" json:"account"`
	PlaceID           uint                 `gorm:"index;primaryKey" json:"placeId"`
	Place             *Place               `gorm:"foreignKey:PlaceID" json:"place,omitempty"`
	AmbassadorRequest bool                 `json:"ambassadorRequest"`
	Type              AccountPlaceRoleType `gorm:"not null" json:"type"`
	CreatedAt         time.Time            `json:"-"`
	UpdatedAt         time.Time            `json:"-"`
	DeletedAt         *time.Time           `gorm:"index" json:"-,omitempty"`
}
