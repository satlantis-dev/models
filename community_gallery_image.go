package models

import "time"

type CommunityGalleryImage struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CommunityID uint      `gorm:"not null;index" json:"communityId"`
	Url         string    `gorm:"not null;unique" json:"url"`
	Caption     *string   `json:"caption"`
	CreatedAt   time.Time `json:"createdAt"`
}
