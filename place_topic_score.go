package models

import (
	"time"
)

type PlaceTopicScore struct {
	CategoryID uint      `gorm:"primaryKey;index;autoIncrement:false" json:"categoryId"`
	PlaceID    uint      `gorm:"primaryKey;autoIncrement:false" json:"placeId"`
	Score      float64   `json:"score"`
	UpdatedAt  time.Time `json:"updatedAt"`
	TopicID    uint      `gorm:"primaryKey;autoIncrement:false" json:"topicId"`
	Topic      Topic     `json:"topic"`
	UserNumber uint      `json:"userNumber"`
	UserScore  float64   `json:"userScore"`
}
