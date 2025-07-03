package models

type GoogleType struct {
	Name         string        `gorm:"primaryKey" json:"name"`
	OSMStdTags   JSONBMapSlice `gorm:"type:jsonb" json:"osmStdTags"`
	OSMExtraTags JSONBMapSlice `gorm:"type:jsonb" json:"osmExtraTags"`
}

func (GoogleType) TableName() string {
	return "google_types"
}
