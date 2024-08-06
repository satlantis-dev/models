package database

import (
	"time"
)

type PlaceMetric struct {
	DataPoints uint      `json:"-"`
	PlaceID    uint      `gorm:"primaryKey" json:"-"`
	MetricID   uint      `gorm:"primaryKey" json:"-"`
	Metric     MetricDTO `json:"metric"`
	UpdatedAt  time.Time `json:"-"`
	Value      float64   `json:"value"`
	ValueStr   string    `json:"valueStr"`
	Score      float64   `json:"score"`
}

type PlaceMetricDTO struct {
	MetricID uint      `json:"metricId"`
	Metric   MetricDTO `json:"metric"`
	Value    float64   `json:"value"`
	ValueStr string    `json:"valueStr"`
}

func (PlaceMetricDTO) TableName() string {
	return "place_metrics"
}
