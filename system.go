package models

type SystemAppVersion struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Platform   string `json:"platform" gorm:"type:varchar(10);not null;uniqueIndex"`
	MinVersion string `json:"minVersion" gorm:"type:varchar(20);not null"`
}
