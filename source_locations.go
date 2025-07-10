package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// OSM

type OSMType string

const (
	OSMTypeNode     OSMType = "node"
	OSMTypeRelation OSMType = "relation"
	OSMTypeWay      OSMType = "way"
)

// Tags

type JSONBMapSlice []map[string]string

func (j *JSONBMapSlice) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(b, j)
}

func (j JSONBMapSlice) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Reports

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

// Reviews

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

// Photos

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

// SourceLocationsOsm  [TEMP STAGING ONLY]

type SourceLocationsOsm struct {
	OSMId              uint             `gorm:"primaryKey;type:uint" json:"osmId"`
	OSMType            OSMType          `gorm:"primaryKey;type:string" json:"osmType"`
	OSMRef             string           `gorm:"uniqueIndex" json:"osmRef"`
	Name               string           `json:"name"`
	Lat                float64          `json:"lat"`
	Lng                float64          `json:"lng"`
	OSMStdTags         JSONBMapSlice    `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags       JSONBMapSlice    `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags      JSONBMapSlice    `gorm:"type:jsonb" json:"satlantisTags"`
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

// SourceLocationsAll

type SourceLocationsAll struct {
	GoogleId           string           `gorm:"primaryKey;index" json:"googleId"`
	OSMRef             string           `json:"osmRef"`
	Name               string           `json:"name"`
	Lat                float64          `json:"lat"`
	Lng                float64          `json:"lng"`
	OSMStdTags         JSONBMapSlice    `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags       JSONBMapSlice    `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags      JSONBMapSlice    `gorm:"type:jsonb" json:"satlantisTags"`
	OSMPlaceRef        string           `json:"osmPlaceRef"`
	OSMPlaceName       string           `json:"osmPlaceName"`
	Source             string           `gorm:"not null" json:"source"`
	OSMDetails         string           `gorm:"type:jsonb" json:"osmDetails"`
	GoogleDetails      string           `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl     string           `json:"googlePhotoUrl"`
	TripadvisorId      *uint            `json:"tripadvisorId"`
	TripadvisorDetails string           `gorm:"type:jsonb" json:"tripadvisorDetails"`
	Reports            Reports          `gorm:"type:jsonb" json:"reports"`
	Reviews            Reviews          `gorm:"type:jsonb" json:"reviews"`
	ReviewSummary      string           `json:"reviewSummary"`
	ReviewHighlights   ReviewHighlights `gorm:"type:jsonb;serializer:json" json:"reviewHighlights"`
	Bio                string           `json:"bio"`
	Hook               string           `gorm:"size:70" json:"hook"`
	Photos             Photos           `gorm:"type:jsonb" json:"photos"`
	Eligible           bool             `json:"eligible"`
	UpdatedAt          time.Time        `json:"-"`
}

func (SourceLocationsAll) TableName() string {
	return "source_locations_all"
}

type SourceLocationDTO struct {
	GoogleId     string  `gorm:"primaryKey;index" json:"googleId"`
	OSMRef       string  `json:"osmRef"`
	Name         string  `json:"name"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	OSMPlaceRef  string  `json:"osmPlaceRef"`
	OSMPlaceName string  `json:"osmPlaceName"`
}

func (sl SourceLocationsAll) ToDTO(db *gorm.DB) (*SourceLocationDTO, error) {
	return &SourceLocationDTO{
		GoogleId:     sl.GoogleId,
		OSMRef:       sl.OSMRef,
		Name:         sl.Name,
		Lat:          sl.Lat,
		Lng:          sl.Lng,
		OSMPlaceRef:  sl.OSMPlaceRef,
		OSMPlaceName: sl.OSMPlaceName,
	}, nil
}
