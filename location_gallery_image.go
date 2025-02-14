package models

import "time"

type ImageCategory string

const (
	ImageCategoryGeneral  ImageCategory = "general"
	ImageCategoryMenu     ImageCategory = "menu"
	ImageCategoryExterior ImageCategory = "exterior"
	ImageCategoryInterior ImageCategory = "interior"
	ImageCategoryFood     ImageCategory = "food"
)

type LocationGalleryImage struct {
	ID         uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	LocationID uint          `gorm:"index" json:"locationId"`
	Url        string        `json:"url"`
	Caption    *string       `json:"caption"`
	Category   ImageCategory `gorm:"type:string;default:'general'" json:"category"`
	Source     string        `json:"source"`
	CreatedAt  time.Time     `json:"-"`
}
