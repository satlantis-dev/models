package models

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseModels = []interface{}{
	Account{},
	AccountPlaceRole{},
	AuthenticationDetail{},
	CalendarEventAccouncement{},
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
		if strings.Contains(err.Error(), "failed to connect to") {
			// Log the error for internal tracking
			log.Println("DB connection error: ", err)

			// Replace sensitive information and use specific error code for client so we know what happened
			return nil, errors.New("internal Server Error - 25874")
		}

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
