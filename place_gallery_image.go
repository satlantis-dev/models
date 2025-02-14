package models

import "time"

type PlaceGalleryImage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PlaceID   uint      `gorm:"index" json:"placeId"`
	Url       string    `json:"url"`
	Caption   *string   `json:"caption"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"-"`
}
