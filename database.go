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
	AccountPlaceRole{},
	AuthenticationDetail{},
	CalendarEvent{},
	CalendarEventRSVP{},
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
	SocialMedia{},
	SourceCountryStats{},
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
