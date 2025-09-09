package models

import "time"

type InterestImage struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	InterestID uint      `gorm:"not null;index" json:"interestId"`
	Interest   *Interest `gorm:"foreignKey:InterestID;references:ID;constraint:OnDelete:CASCADE;" json:"interest,omitempty"`
	Url        string    `gorm:"not null;unique" json:"url"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"_"`
}
