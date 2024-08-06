package database

import "time"

type SourceFacts struct {
	PlaceOSM    string    `gorm:"primaryKey;" json:"placeOsm"`
	PlaceName   string    `json:"placeName"`
	CountryName string    `json:"countryName"`
	Facts       string    `gorm:"type:jsonb" json:"facts"`
	UpdatedOn   time.Time `json:"updatedOn"`
}

func (SourceFacts) TableName() string {
	return "source_facts"
}
