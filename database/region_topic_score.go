package database

import (
	"time"
)

type RegionTopicScore struct {
	CategoryID uint      `gorm:"primaryKey;index;autoIncrement:false" json:"categoryId"`
	RegionID   uint      `gorm:"primaryKey;autoIncrement:false" json:"regionId"`
	Score      float64   `json:"score"`
	UpdatedAt  time.Time `json:"updatedAt"`
	TopicID    uint      `gorm:"primaryKey;autoIncrement:false" json:"topicId"`
	Topic      Topic     `json:"topic"`
	UserNumber uint      `json:"userNumber"`
	UserScore  float64   `json:"userScore"`
}
