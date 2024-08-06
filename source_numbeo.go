package models

import "time"

// SourceNumbeo with Response as JSON object
type SourceNumbeo struct {
	PlaceOSM    string    `gorm:"primaryKey;" json:"placeOsm"`
	PlaceName   string    `json:"placeName"`
	CountryName string    `json:"countryName"`
	Domain      string    `gorm:"primaryKey;size:10" json:"domain"`
	Response    string    `gorm:"type:jsonb" json:"response"`
	UpdatedOn   time.Time `json:"updatedOn"`
}

func (SourceNumbeo) TableName() string {
	return "source_numbeo"
}
