package models

import "time"

type LocationGalleryImage struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	LocationID uint      `json:"locationId"`
	Location   *Location `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	Url        string    `json:"url"`
	Caption    *string   `json:"caption"`
	Category   *string   `json:"category"`
	Source     string    `json:"source"`
	CreatedAt  time.Time `json:"createdAt"`
}
