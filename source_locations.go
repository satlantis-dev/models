package models

import "time"

type OSMType string

const (
	OSMTypeNode     OSMType = "node"
	OSMTypeRelation OSMType = "relation"
	OSMTypeWay      OSMType = "way"
)

// SourceLocationsOsm with Response as JSON object
type SourceLocationsOsm struct {
	OSMId          uint      `gorm:"primaryKey;type:uint" json:"osmId"`
	OSMType        OSMType   `gorm:"primaryKey;type:string" json:"osmType"`
	OSMRef         string    `gorm:"uniqueIndex" json:"osmRef"`
	Name           string    `json:"name"`
	Lat            float64   `json:"lat"`
	Lng            float64   `json:"lng"`
	OSMStdTags     string    `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags   string    `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags  string    `gorm:"type:jsonb" json:"satlantisTags"`
	OSMDetails     string    `gorm:"type:jsonb" json:"osmDetails"`
	OSMPlaceRef    string    `json:"osmPlaceRef"`
	OSMPlaceName   string    `json:"osmPlaceName"`
	Source         string    `gorm:"default:osm" json:"source"`
	GoogleId       string    `json:"googleId"`
	GoogleDetails  string    `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl string    `json:"googlePhotoUrl"`
	UpdatedOn      time.Time `json:"updatedOn"`
	Eligible       bool      `json:"eligible"`
}

func (SourceLocationsOsm) TableName() string {
	return "source_locations_osm"
}

// SourceLocationsExtra with Response as JSON object
type SourceLocationsExtra struct {
	OSMId          uint      `gorm:"primaryKey;type:bigserial" json:"osmId"`
	OSMType        OSMType   `gorm:"type:string;default:node" json:"osmType"`
	OSMRef         string    `gorm:"uniqueIndex" json:"osmRef"`
	Name           string    `json:"name"`
	Lat            float64   `json:"lat"`
	Lng            float64   `json:"lng"`
	OSMStdTags     string    `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags   string    `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags  string    `gorm:"type:jsonb" json:"satlantisTags"`
	OSMDetails     string    `gorm:"type:jsonb" json:"osmDetails"`
	OSMPlaceRef    string    `json:"osmPlaceRef"`
	OSMPlaceName   string    `json:"osmPlaceName"`
	Source         string    `json:"source"`
	GoogleId       string    `json:"googleId"`
	GoogleDetails  string    `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl string    `json:"googlePhotoUrl"`
	UpdatedOn      time.Time `json:"updatedOn"`
	Eligible       bool      `json:"eligible"`
}

func (SourceLocationsExtra) TableName() string {
	return "source_locations_extra"
}

type SourceLocations struct {
	OSMId          uint      `json:"osmId"`
	OSMType        OSMType   `json:"osmType"`
	OSMRef         string    `gorm:"uniqueIndex" json:"osmRef"`
	Name           string    `json:"name"`
	Lat            float64   `json:"lat"`
	Lng            float64   `json:"lng"`
	OSMStdTags     string    `json:"osmStdTags"`
	OSMExtraTags   string    `json:"osmExtraTags"`
	SatlantisTags  string    `json:"satlantisTags"`
	OSMDetails     string    `json:"osmDetails"`
	OSMPlaceRef    string    `json:"osmPlaceRef"`
	OSMPlaceName   string    `json:"osmPlaceName"`
	Source         string    `json:"source"`
	GoogleId       string    `json:"googleId"`
	GoogleDetails  string    `json:"googleDetails"`
	GooglePhotoUrl string    `json:"googlePhotoUrl"`
	UpdatedOn      time.Time `json:"updatedOn"`
	Eligible       bool      `json:"eligible"`
}
