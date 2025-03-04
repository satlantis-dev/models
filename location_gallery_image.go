package models

import "time"

type ImageCategory string

const (
	ImageCategoryGeneral       ImageCategory = "general"
	ImageCategoryExterior      ImageCategory = "exterior"
	ImageCategoryInterior      ImageCategory = "interior"
	ImageCategoryAmenities     ImageCategory = "amenities"
	ImageCategoryFoodAndDrinks ImageCategory = "foodandrinks"
	ImageCategoryMenu          ImageCategory = "menu"
)

type Photo struct {
	LocationID uint          `gorm:"not null;index" json:"locationId"`
	Url        string        `gorm:"not null;unique" json:"url"`
	Caption    *string       `json:"caption"`
	Category   ImageCategory `gorm:"type:string;default:'general'" json:"category"`
	Highlight  bool          `gorm:"default:false" json:"highlight"`
	Source     string        `json:"source"`
}

type LocationGalleryImage struct {
	ID         uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	LocationID uint          `gorm:"not null;index" json:"locationId"`
	Url        string        `gorm:"not null;unique" json:"url"`
	Caption    *string       `json:"caption"`
	Category   ImageCategory `gorm:"type:string;default:'general'" json:"category"`
	Highlight  bool          `gorm:"default:false" json:"highlight"`
	Source     string        `json:"source"`
	CreatedAt  time.Time     `json:"-"`
}
