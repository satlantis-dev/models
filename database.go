package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseModels = []interface{}{
	Account{},
	AccountNoteTracker{},
	AccountPlaceRole{},
	AuthenticationDetail{},
	CalendarEventAnnouncement{},
	CalendarEventNote{},
	CalendarEventRSVP{},
	CalendarEvent{},
	Category{},
	ChatMembership{},
	Continent{},
	Country{},
	Currency{},
	Event{},
	Experience{},
	Follow{},
	Interest{},
	Location{},
	LocationAccount{},
	LocationNote{},
	LocationTag{},
	Metric{},
	MetricSource{},
	Note{},
	NoteClosure{},
	NoteRanking{},
	Notification{},
	Place{},
	PlaceCalendarEvent{},
	PlaceCategoryScore{},
	PlaceClosure{},
	PlaceFoundationScore{},
	PlaceMetric{},
	PlaceNote{},
	PlacePublicChatChannel{},
	PlaceTopicScore{},
	PublicChatChannel{},
	PublicChatMembership{},
	PublicChatMessage{},
	Reaction{},
	RegionCategoryScore{},
	RegionMetric{},
	RegionTopicScore{},
	Region{},
	Relay{},
	SocialMedia{},
	SourceBots{},
	SourceCountryStats{},
	SourceGoogleAnalytics{},
	SourceNumbeo{},
	SourceRuleoflaw{},
	SourceSpeedtest{},
	SourceFacts{},
	SourceSurvey{},
	SourceLocationsOsm{},
	SourceLocationsExtra{},
	Tag{},
	TeamMember{},
	Topic{},
	VerificationToken{},
	Weather{},
	RestrictedUsername{},
	AccountAnalytics{},
	MerchantAnalytics{},
	EventAnalytics{},
	EngagementAnalytics{},
}

func SetupDatabase(dbHost, dbUser, dbPassword, dbName string) (*gorm.DB, error) {
	var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sqlDB: %v", err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	return db, nil
}
