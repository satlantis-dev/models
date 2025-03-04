package models

import "time"

type PlaceGalleryImage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PlaceID   uint      `gorm:"not null;index" json:"placeId"`
	Url       string    `gorm:"not null;unique" json:"url"`
	Caption   *string   `json:"caption"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}
