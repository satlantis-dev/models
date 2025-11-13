package models

import (
	"time"
)

type PlaceTopicScore struct {
	CategoryID uint      `gorm:"primaryKey;index;autoIncrement:false" json:"categoryId"`
	PlaceID    uint      `gorm:"primaryKey;autoIncrement:false" json:"placeId"`
	Place      Place     `gorm:"constraint:OnDelete:CASCADE" json:"-"`
	Score      float64   `json:"score"`
	UpdatedAt  time.Time `json:"updatedAt"`
	TopicID    uint      `gorm:"primaryKey;autoIncrement:false" json:"topicId"`
	Topic      Topic     `gorm:"foreignKey:TopicID;constraint:OnDelete:CASCADE" json:"topic"`
	UserNumber uint      `json:"userNumber"`
	UserScore  float64   `json:"userScore"`
}
