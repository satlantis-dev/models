package database

import "time"

type PlaceFoundationScore struct {
	PlaceID   uint      `gorm:"primaryKey" json:"cityId"`
	Score     float64   `json:"score"`
	UpdatedAt time.Time `json:"updatedAt"`
	TopicID   uint      `gorm:"primaryKey" json:"topicId"`
}
