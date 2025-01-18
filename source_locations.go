package models

import "time"

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

type Review struct {
	Source string    `json:"source"`
	Author string    `json:"author"`
	Rating string    `json:"rating"`
	Time   time.Time `json:"time"`
	Text   string    `json:"text"`
}

// SourceLocationsOsm with Response as JSON object
type SourceLocationsOsm struct {
	OSMId              uint      `gorm:"primaryKey;type:uint" json:"osmId"`
	OSMType            OSMType   `gorm:"primaryKey;type:string" json:"osmType"`
	OSMRef             string    `gorm:"uniqueIndex" json:"osmRef"`
	Name               string    `json:"name"`
	Lat                float64   `json:"lat"`
	Lng                float64   `json:"lng"`
	OSMStdTags         string    `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags       string    `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags      string    `gorm:"type:jsonb" json:"satlantisTags"`
	OSMDetails         string    `gorm:"type:jsonb" json:"osmDetails"`
	OSMPlaceRef        string    `json:"osmPlaceRef"`
	OSMPlaceName       string    `json:"osmPlaceName"`
	Source             string    `gorm:"default:osm" json:"source"`
	GoogleId           string    `json:"googleId"`
	GoogleDetails      string    `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl     string    `json:"googlePhotoUrl"`
	TripadvisorId      *uint     `json:"tripadvisorId"`
	TripadvisorDetails string    `gorm:"type:jsonb" json:"tripadvisorDetails"`
	UpdatedOn          time.Time `json:"updatedOn"`
	Eligible           bool      `json:"eligible"`
	Reports            []Report  `gorm:"type:jsonb" json:"reports"`
	ReviewSummary      string    `json:"reviewSummary"`
	Reviews            []Review  `gorm:"type:jsonb" json:"reviews"`
}

func (SourceLocationsOsm) TableName() string {
	return "source_locations_osm"
}

// SourceLocationsExtra with Response as JSON object
type SourceLocationsExtra struct {
	OSMId              uint      `gorm:"primaryKey;type:bigserial" json:"osmId"`
	OSMType            OSMType   `gorm:"type:string;default:node" json:"osmType"`
	OSMRef             string    `gorm:"uniqueIndex" json:"osmRef"`
	Name               string    `json:"name"`
	Lat                float64   `json:"lat"`
	Lng                float64   `json:"lng"`
	OSMStdTags         string    `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags       string    `gorm:"type:jsonb" json:"osmExtraTags"`
	SatlantisTags      string    `gorm:"type:jsonb" json:"satlantisTags"`
	OSMDetails         string    `gorm:"type:jsonb" json:"osmDetails"`
	OSMPlaceRef        string    `json:"osmPlaceRef"`
	OSMPlaceName       string    `json:"osmPlaceName"`
	Source             string    `gorm:"not null" json:"source"`
	GoogleId           string    `gorm:"uniqueIndex" json:"googleId"`
	GoogleDetails      string    `gorm:"type:jsonb" json:"googleDetails"`
	GooglePhotoUrl     string    `json:"googlePhotoUrl"`
	TripadvisorId      *uint     `json:"tripadvisorId"`
	TripadvisorDetails string    `gorm:"type:jsonb" json:"tripadvisorDetails"`
	UpdatedOn          time.Time `json:"updatedOn"`
	Eligible           bool      `json:"eligible"`
	Reports            []Report  `gorm:"type:jsonb" json:"reports"`
	ReviewSummary      string    `json:"reviewSummary"`
	Reviews            []Review  `gorm:"type:jsonb" json:"reviews"`
}

func (SourceLocationsExtra) TableName() string {
	return "source_locations_extra"
}

type SourceLocations struct {
	OSMId              uint      `json:"osmId"`
	OSMType            OSMType   `json:"osmType"`
	OSMRef             string    `gorm:"uniqueIndex" json:"osmRef"`
	Name               string    `json:"name"`
	Lat                float64   `json:"lat"`
	Lng                float64   `json:"lng"`
	OSMStdTags         string    `json:"osmStdTags"`
	OSMExtraTags       string    `json:"osmExtraTags"`
	SatlantisTags      string    `json:"satlantisTags"`
	OSMDetails         string    `json:"osmDetails"`
	OSMPlaceRef        string    `json:"osmPlaceRef"`
	OSMPlaceName       string    `json:"osmPlaceName"`
	Source             string    `json:"source"`
	GoogleId           string    `json:"googleId"`
	GoogleDetails      string    `json:"googleDetails"`
	GooglePhotoUrl     string    `json:"googlePhotoUrl"`
	TripadvisorId      *uint     `json:"tripadvisorId"`
	TripadvisorDetails string    `json:"tripadvisorDetails"`
	UpdatedOn          time.Time `json:"updatedOn"`
	Eligible           bool      `json:"eligible"`
	Reports            []Report  `gorm:"type:jsonb" json:"reports"`
	ReviewSummary      string    `json:"reviewSummary"`
	Reviews            []Review  `gorm:"type:jsonb" json:"reviews"`
}
