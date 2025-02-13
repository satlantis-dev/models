package models

import "time"

type PlaceGalleryImage struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PlaceID   uint      `gorm:"index" json:"placeId"`
	Place     *Place    `gorm:"foreignKey:PlaceID" json:"place,omitempty"`
	Url       string    `json:"url"`
	Caption   *string   `json:"caption"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"createdAt"`
}
