package models

import (
	"time"
)

type AccountAnalytics struct {
	Date             time.Time `gorm:"primaryKey;type:date" json:"date"`
	Signups          int       `json:"signups"`
	SignupsSatlantis int       `json:"signupsSatlantis"`
	Visitors         int       `json:"visitors"`
}

type MerchantAnalytics struct {
	Date           time.Time `gorm:"primaryKey;type:date" json:"date"`
	PlaceID        uint      `gorm:"primaryKey" json:"placeId"`
	Place          Place     `gorm:"foreignKey:PlaceID;references:ID"`
	Locations      int       `json:"locations"`
	TotalLocations int       `json:"totalLocations"`
	Claimed        int       `json:"claimed"`
	TotalClaimed   int       `json:"totalClaimed"`
}

type EventAnalytics struct {
	Date            time.Time `gorm:"primaryKey;type:date" json:"date"`
	PlaceID         uint      `gorm:"primaryKey" json:"placeId"`
	Place           Place     `gorm:"foreignKey:PlaceID;references:ID"`
	Events          int       `json:"events"`
	EventsSatlantis int       `json:"eventsSatlantis"`
	UpcomingEvents  int       `json:"upcomingEvents"`
}

type EngagementAnalytics struct {
	Date           time.Time `gorm:"primaryKey;type:date" json:"date"`
	MediaNoteViews int       `json:"mediaNoteViews"`
	MediaNotes     int       `json:"mediaNotes"`
	Reactions      int       `json:"reactions"`
	Comments       int       `json:"comments"`
}

type PlaceAnalytics struct {
	Date                time.Time `gorm:"primaryKey;type:date" json:"date"`
	PlaceID             uint      `gorm:"primaryKey" json:"placeId"`
	Place               Place     `gorm:"foreignKey:PlaceID;references:ID"`
	Views               int       `json:"views"`
	Followers           int       `json:"followers"`
	Unfollowers         int       `json:"unfollowers"`
	NetFollowers        int       `json:"netFollowers"`
	TotalFollowers      int       `json:"totalFollowers"`
	MediaNotes          int       `json:"mediaNotes"`
	MediaNotesSatlantis int       `json:"mediaNotesSatlantis"`
}
