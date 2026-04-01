package models

import "time"

type CommunityGalleryImage struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID uint      `gorm:"not null;index" json:"communityId"`
	Url         string    `gorm:"not null;unique" json:"url"`
	Rank        int       `gorm:"not null;default:0" json:"rank"`
	CreatedAt   time.Time `json:"createdAt"`
}
