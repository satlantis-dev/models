package models

type GoogleType struct {
	Name        string             `gorm:"primaryKey" json:"name"`
	OSMStdTag   *map[string]string `gorm:"type:jsonb" json:"osmStdTag"`
	OSMExtraTag *map[string]string `gorm:"type:jsonb" json:"osmExtraTag"`
}

func (GoogleType) TableName() string {
	return "google_types"
}
