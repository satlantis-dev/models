package models

import (
	"time"
)

type SourceGoogleAnalytics struct {
	Date                      time.Time `gorm:"primaryKey;type:date" json:"date"`
	Active1DayUsers           int       `json:"active1DayUsers"`
	Active28DayUsers          int       `json:"active28DayUsers"`
	Active7DayUsers           int       `json:"active7DayUsers"`
	AverageSessionDuration    float64   `json:"averageSessionDuration"`
	EngagedSessions           int       `json:"engagedSessions"`
	EngagementRate            float64   `json:"engagementRate"`
	EventCount                int       `json:"eventCount"`
	EventCountPerUser         float64   `json:"eventCountPerUser"`
	EventsPerSession          float64   `json:"eventsPerSession"`
	NewUsers                  int       `json:"newUsers"`
	ScreenPageViews           int       `json:"screenPageViews"`
	ScreenPageViewsPerSession float64   `json:"screenPageViewsPerSession"`
	ScreenPageViewsPerUser    float64   `json:"screenPageViewsPerUser"`
	Sessions                  int       `json:"sessions"`
	SessionsPerUser           float64   `json:"sessionsPerUser"`
	UserEngagementDuration    float64   `json:"userEngagementDuration"`
	TopPages                  string    `gorm:"type:jsonb" json:"topPages"`
}