package models

import (
	"github.com/lib/pq"
)

type MetricFormat int

const (
	ONE_TO_FIVE MetricFormat = iota + 1
	ZERO_TO_ONEHUNDRED_HIGH_BETTER
	ZERO_TO_ONEHUNDRED_LOW_BETTER
	AMOUNT_LOC_CURR
	NUMERIC_HIGH_BETTER
	NUMERIC_LOW_BETTER
	YES_NO
	NON_NUMERIC
)

type Metric struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	CategoryID      uint           `gorm:"index" json:"categoryId"`
	Category        Category       `json:"category"`
	Description     string         `gorm:"type:text" json:"description"`
	Format          MetricFormat   `json:"type"`
	MetricSourceID  uint           `gorm:"index" json:"metric_sourceId"`
	MetricSource    MetricSource   `json:"metricSource"`
	Name            string         `gorm:"type:text" json:"name"`
	Prompt          string         `gorm:"type:text" json:"prompt"`
	Slug            string         `gorm:"type:text" json:"slug"`
	MetricSourceRef string         `json:"metricSourceRef"`
	Suffix          string         `gorm:"type:text" json:"suffix"`
	TopicID         uint           `gorm:"index" json:"topicId"`
	Topic           Topic          `json:"topic"`
	Weight          float64        `json:"weight"`
	IsScorable      bool           `json:"isScorable"`
	Tags            pq.StringArray `gorm:"type:varchar[]" json:"tags"`
	Order           *uint          `json:"order"`
}

type MetricDTO struct {
	ID          uint     `json:"id"`
	CategoryID  uint     `json:"categoryId"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Prompt      string   `json:"prompt"`
	Slug        string   `json:"slug"`
	Suffix      string   `json:"suffix"`
	Tags        string   `json:"tags"`
	TopicID     uint     `json:"topicId"`
	Topic       Topic    `json:"topic"`
}

func (MetricDTO) TableName() string {
	return "metrics"
}
