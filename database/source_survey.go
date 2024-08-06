package database

import "time"

type SourceSurvey struct {
	PlaceOSM    string    `gorm:"primaryKey;" json:"placeOsm"`
	PlaceName   string    `json:"placeName"`
	CountryName string    `json:"countryName"`
	Entries     string    `gorm:"type:jsonb" json:"entries"`
	UpdatedOn   time.Time `json:"updatedOn"`
}

func (SourceSurvey) TableName() string {
	return "source_survey"
}
