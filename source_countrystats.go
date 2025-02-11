package models

import "time"

// SourceGdp with Response as JSON object
type SourceCountryStats struct {
	CountryCode            string    `gorm:"index;primaryKey;type:char(3)" json:"countryCode"`
	RealGdpGrowth          float64   `json:"realGdpGrowth"`
	GdpPerCapitaPpp        float64   `json:"gdpPerCapitaPpp"`
	ConsumerPriceInflation float64   `json:"consumerPriceInflation"`
	HappinessIndex         float64   `json:"happinessIndex"`
	BusinessIndex          float64   `json:"businessIndex"`
	UpdatedOn              time.Time `json:"updatedOn"`
}

func (SourceCountryStats) TableName() string {
	return "source_countrystats"
}
