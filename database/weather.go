package database

import (
	"time"
)

type Weather struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-,omitempty"`
	PlaceID   uint       `gorm:"index" json:"placeId"`
	Humidity  int        `json:"humidity"`
	Pressure  int        `json:"pressure"`
	Temp      float64    `json:"temp"`
}
