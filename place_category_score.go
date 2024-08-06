package models

import (
	"time"
)

type PlaceCategoryScore struct {
	CategoryID  uint              `gorm:"primaryKey;autoIncrement:false" json:"categoryId"`
	Category    Category          `json:"category"`
	PlaceID     uint              `gorm:"primaryKey;autoIncrement:false" json:"-"`
	Score       float64           `json:"score"`
	TopicScores []PlaceTopicScore `gorm:"foreignKey:CategoryID,PlaceID;references:CategoryID,PlaceID" json:"topicScores"`
	UpdatedAt   time.Time         `json:"-"`
}
