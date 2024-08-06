package database

import (
	"time"
)

type RegionCategoryScore struct {
	CategoryID  uint               `gorm:"primaryKey;autoIncrement:false" json:"categoryId"`
	Category    Category           `json:"category"`
	RegionID    uint               `gorm:"primaryKey;autoIncrement:false" json:"-"`
	Score       float64            `json:"score"`
	TopicScores []RegionTopicScore `gorm:"foreignKey:CategoryID,RegionID;references:CategoryID,RegionID" json:"topicScores"`
	UpdatedAt   time.Time          `json:"-"`
}
