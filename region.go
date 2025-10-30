package models

import (
	"time"
)

type Region struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Code      string     `gorm:"type:text" json:"code"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-,omitempty"`
	CountryID uint       `gorm:"index" json:"countryId"`
	Country   *Country   `json:"country,omitempty"`
	Name      string     `gorm:"index;type:text" json:"name"`
	OSMID     *uint      `json:"osmId"`
	OSMLevel  string     `json:"osmLevel"`
	OSMType   OSMType    `json:"osmType"`
	OSMRef    string     `gorm:"uniqueIndex" json:"osmRef"`
	Places    *[]Place   `gorm:"foreignKey:RegionID" json:"places,omitempty"`
}
