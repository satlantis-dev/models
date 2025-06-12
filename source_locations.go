package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type OSMType string

const (
	OSMTypeNode     OSMType = "node"
	OSMTypeRelation OSMType = "relation"
	OSMTypeWay      OSMType = "way"
)

type ReportType string

const (
	ReportTypeWrongInfo   ReportType = "wrongInfo"
	ReportTypeMissingInfo ReportType = "missingInfo"
	ReportTypeDuplicate   ReportType = "duplicate"
	ReportTypeClosed      ReportType = "closed"
	ReportTypeAdd         ReportType = "add"
	ReportTypeRemove      ReportType = "remove"
)

type Report struct {
	Type      ReportType `json:"reportType"`
	AccountId uint       `json:"accountId"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"createdAt"`
}

type Reports []Report

func (r *Reports) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, r)
}

func (r Reports) Value() (driver.Value, error) {
	return json.Marshal(r)
}

type Review struct {
	Source string    `json:"source"`
	Id     string    `json:"id"`
	Author string    `json:"author"`
	Rating float64   `json:"rating"`
	Time   time.Time `json:"time"`
	Text   string    `json:"text"`
}

type Reviews []Review

func (r *Reviews) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, r)
}

func (r Reviews) Value() (driver.Value, error) {
	return json.Marshal(r)
}

type Photo struct {
	Url      string `gorm:"not null;unique" json:"url"`
	SourceID string `gorm:"not null;unique" json:"id"`
	Source   string `json:"source"`
	Width    *int   `json:"width"`
	Height   *int   `json:"height"`
}

type Photos []Photo

func (p *Photos) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, p)
}

func (p Photos) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// SourceLocationsOsm with Response as JSON object
type SourceLocationsOsm struct {
	OSMId              uint             `gorm:"primaryKey;type:uint" json:"osmId"`
	OSMType            OSMType          `gorm:"primaryKey;type:string" json:"osmType"`
	OSMRef             string           `gorm:"uniqueIndex" json:"osmRef"`
	Name               string           `json:"name"`
	Lat                float64          `json:"lat"`
	Lng                float64          `json:"lng"`
	OSMStdTags         string           `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags       string           `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags      string           `gorm:"type:jsonb" json:"satlantisTags"`
	OSMDetails         string           `gorm:"type:jsonb" json:"osmDetails"`
	OSMPlaceRef        string           `json:"osmPlaceRef"`
	OSMPlaceName       string           `json:"osmPlaceName"`
	Source             string           `gorm:"default:osm" json:"source"`
	GoogleId           string           `gorm:"uniqueIndex" json:"googleId"`
	GoogleDetails      string           `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl     string           `json:"googlePhotoUrl"`
	TripadvisorId      *uint            `json:"tripadvisorId"`
	TripadvisorDetails string           `gorm:"type:jsonb" json:"tripadvisorDetails"`
	UpdatedOn          time.Time        `json:"updatedOn"`
	Eligible           bool             `json:"eligible"`
	Reports            Reports          `gorm:"type:jsonb" json:"reports"`
	Reviews            Reviews          `gorm:"type:jsonb" json:"reviews"`
	ReviewSummary      string           `json:"reviewSummary"`
	ReviewHighlights   ReviewHighlights `gorm:"type:jsonb;serializer:json" json:"reviewHighlights"`
	Bio                string           `json:"bio"`
	Hook               string           `gorm:"size:70" json:"hook"`
	Photos             Photos           `gorm:"type:jsonb" json:"photos"`
}

func (SourceLocationsOsm) TableName() string {
	return "source_locations_osm"
}

// SourceLocationsExtra with Response as JSON object
type SourceLocationsExtra struct {
	OSMId              uint             `gorm:"primaryKey;type:bigserial" json:"osmId"`
	OSMType            OSMType          `gorm:"type:string;default:node" json:"osmType"`
	OSMRef             string           `gorm:"uniqueIndex" json:"osmRef"`
	Name               string           `json:"name"`
	Lat                float64          `json:"lat"`
	Lng                float64          `json:"lng"`
	OSMStdTags         string           `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags       string           `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags      string           `gorm:"type:jsonb" json:"satlantisTags"`
	OSMDetails         string           `gorm:"type:jsonb" json:"osmDetails"`
	OSMPlaceRef        string           `json:"osmPlaceRef"`
	OSMPlaceName       string           `json:"osmPlaceName"`
	Source             string           `gorm:"not null" json:"source"`
	GoogleId           string           `gorm:"uniqueIndex" json:"googleId"`
	GoogleDetails      string           `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl     string           `json:"googlePhotoUrl"`
	TripadvisorId      *uint            `json:"tripadvisorId"`
	TripadvisorDetails string           `gorm:"type:jsonb" json:"tripadvisorDetails"`
	UpdatedOn          time.Time        `json:"updatedOn"`
	Eligible           bool             `json:"eligible"`
	Reports            Reports          `gorm:"type:jsonb" json:"reports"`
	Reviews            Reviews          `gorm:"type:jsonb" json:"reviews"`
	ReviewSummary      string           `json:"reviewSummary"`
	ReviewHighlights   ReviewHighlights `gorm:"type:jsonb;serializer:json" json:"reviewHighlights"`
	Bio                string           `json:"bio"`
	Hook               string           `gorm:"size:70" json:"hook"`
	Photos             Photos           `gorm:"type:jsonb" json:"photos"`
}

func (SourceLocationsExtra) TableName() string {
	return "source_locations_extra"
}

// SourceLocationsAll with Response as JSON object
type SourceLocationsAll struct {
	GoogleId           string              `gorm:"primaryKey;index" json:"googleId"`
	OSMRef             string              `gorm:"unique" json:"osmRef"`
	Name               string              `json:"name"`
	Lat                float64             `json:"lat"`
	Lng                float64             `json:"lng"`
	OSMStdTags         []map[string]string `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags       []map[string]string `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags      []map[string]string `gorm:"type:jsonb" json:"satlantisTags"`
	OSMPlaceRef        string              `json:"osmPlaceRef"`
	OSMPlaceName       string              `json:"osmPlaceName"`
	Source             string              `gorm:"not null" json:"source"`
	OSMDetails         string              `gorm:"type:jsonb" json:"osmDetails"`
	GoogleDetails      string              `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl     string              `json:"googlePhotoUrl"`
	TripadvisorId      *uint               `json:"tripadvisorId"`
	TripadvisorDetails string              `gorm:"type:jsonb" json:"tripadvisorDetails"`
	Reports            Reports             `gorm:"type:jsonb" json:"reports"`
	Reviews            Reviews             `gorm:"type:jsonb" json:"reviews"`
	ReviewSummary      string              `json:"reviewSummary"`
	ReviewHighlights   ReviewHighlights    `gorm:"type:jsonb;serializer:json" json:"reviewHighlights"`
	Bio                string              `json:"bio"`
	Hook               string              `gorm:"size:70" json:"hook"`
	Photos             Photos              `gorm:"type:jsonb" json:"photos"`
	Eligible           bool                `json:"eligible"`
	UpdatedAt          time.Time           `json:"-"`
}

func (SourceLocationsAll) TableName() string {
	return "source_locations_all"
}

type SourceLocations struct {
	OSMId              uint             `json:"osmId"`
	OSMType            OSMType          `json:"osmType"`
	OSMRef             string           `gorm:"uniqueIndex" json:"osmRef"`
	Name               string           `json:"name"`
	Lat                float64          `json:"lat"`
	Lng                float64          `json:"lng"`
	OSMStdTags         string           `json:"osmStdTags"`
	OSMExtraTags       string           `json:"osmExtraTags"`
	SatlantisTags      string           `json:"satlantisTags"`
	OSMDetails         string           `json:"osmDetails"`
	OSMPlaceRef        string           `json:"osmPlaceRef"`
	OSMPlaceName       string           `json:"osmPlaceName"`
	Source             string           `json:"source"`
	GoogleId           string           `json:"googleId"`
	GoogleDetails      string           `json:"googleDetails"`
	GooglePhotoUrl     string           `json:"googlePhotoUrl"`
	TripadvisorId      *uint            `json:"tripadvisorId"`
	TripadvisorDetails string           `json:"tripadvisorDetails"`
	UpdatedOn          time.Time        `json:"updatedOn"`
	Eligible           bool             `json:"eligible"`
	Reports            Reports          `gorm:"type:jsonb" json:"reports"`
	Reviews            Reviews          `gorm:"type:jsonb" json:"reviews"`
	ReviewSummary      string           `json:"reviewSummary"`
	ReviewHighlights   ReviewHighlights `gorm:"type:jsonb;serializer:json" json:"reviewHighlights"`
	Bio                string           `json:"bio"`
	Hook               string           `gorm:"size:70" json:"hook"`
	Photos             Photos           `gorm:"type:jsonb" json:"photos"`
}
