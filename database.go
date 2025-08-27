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
	AccountCollectionRole{},
	AccountNoteTracker{},
	AccountPlaceRole{},
	AccountLocationRole{},
	AccountLocationReview{},
	AuthenticationDetail{},
	Block{},
	Brand{},
	CalendarEventAnnouncement{},
	CalendarEventNote{},
	CalendarEventRSVP{},
	Calendar{},
	CalendarEvent{},
	Category{},
	ChatMembership{},
	Cluster{},
	ClusterInterest{},
	Collection{},
	CollectionLocation{},
	Continent{},
	Country{},
	Currency{},
	Event{},
	Follow{},
	GoogleType{},
	Interest{},
	Location{},
	LocationCategory{},
	LocationClaim{},
	LocationDirectory{},
	LocationGalleryImage{},
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
	PlaceGalleryImage{},
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
	Save{},
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
	SourceLocationsAll{},
	Tag{},
	TeamMember{},
	Topic{},
	VerificationToken{},
	Weather{},
	RestrictedUsername{},
	AccountAnalytics{},
	EngagementAnalytics{},
	EventAnalytics{},
	MerchantAnalytics{},
	PlaceAnalytics{},
	Device{},
	Conversation{},
	CalendarEventCohost{},
	ChatMessage{},
	ContentReport{},
	NoteRepost{},
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
