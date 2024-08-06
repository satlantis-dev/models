package database

import (
	"time"
)

type RegionMetric struct {
	DataPoints uint      `json:"-"`
	RegionID   uint      `gorm:"primaryKey" json:"-"`
	MetricID   uint      `gorm:"primaryKey" json:"-"`
	Metric     MetricDTO `json:"metric"`
	UpdatedAt  time.Time `json:"-"`
	Value      float64   `json:"value"`
	ValueStr   string    `json:"valueStr"`
	Score      float64   `json:"score"`
}

type RegionMetricDTO struct {
	MetricID uint      `json:"metricId"`
	Metric   MetricDTO `json:"metric"`
	Value    float64   `json:"value"`
	ValueStr string    `json:"valueStr"`
}

func (RegionMetricDTO) TableName() string {
	return "region_metrics"
}
