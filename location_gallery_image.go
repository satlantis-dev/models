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
	ID         uint          `gorm:"primaryKey" json:"id"`
	LocationID uint          `gorm:"index" json:"locationId"`
	Location   *Location     `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	Url        string        `json:"url"`
	Caption    *string       `json:"caption"`
	Category   ImageCategory `gorm:"type:string,default:general" json:"category"`
	Source     string        `json:"source"`
	CreatedAt  time.Time     `json:"createdAt"`
}
