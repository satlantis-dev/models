package database

import "time"

// SourceRuleoflaw with Response as JSON object
type SourceRuleoflaw struct {
	CountryCode string    `gorm:"primaryKey;type:char(3)" json:"countryCode"`
	Scores      string    `gorm:"type:jsonb" json:"scores"`
	UpdatedOn   time.Time `json:"updatedOn"`
}

func (SourceRuleoflaw) TableName() string {
	return "source_ruleoflaw"
}
