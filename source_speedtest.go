package models

import "time"

// SourceSpeedtest with Response as JSON object
type SourceSpeedtest struct {
	Entity              string    `gorm:"primaryKey;size:100" json:"entity"`
	FixedDownloadSpeed  float64   `json:"fixedDownloadSpeed"`
	FixedUploadSpeed    float64   `json:"fixedUploadSpeed"`
	FixedLatency        float64   `json:"fixedLatency"`
	Level               string    `json:"level"`
	MobileDownloadSpeed float64   `json:"mobileDownloadSpeed"`
	MobileUploadSpeed   float64   `json:"mobileUploadSpeed"`
	MobileLatency       float64   `json:"mobileLatency"`
	UpdatedOn           time.Time `json:"updatedOn"`
}

func (SourceSpeedtest) TableName() string {
	return "source_speedtest"
}
