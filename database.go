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
	AccountAnalytics{},
	AccountCalendarRole{},
	AccountCollectionRole{},
	AccountNoteTracker{},
	AccountLocationRole{},
	AccountLocationReview{},
	AccountPlaceRole{},
	AccountRSVPAnswers{},
	Activity{},
	Block{},
	Brand{},
	Calendar{},
	CalendarEvent{},
	CalendarEventAnnouncement{},
	CalendarEventCohost{},
	CalendarEventNote{},
	CalendarEventRSVP{},
	CalendarEventTicket{},
	CalendarEventTicketType{},
	CalendarEventTicketOrder{},
	CalendarEventTicketOrderItem{},
	CalendarEventTicketOrderPayment{},
	CalendarEventTicketOrderRefund{},
	CalendarEventWallet{},
	CalendarEventWalletWithdrawal{},
	CalendarSubscription{},
	CalendarEventTag{},
	Category{},
	ChatMembership{},
	ChatMessage{},
	Cluster{},
	ClusterInterest{},
	Collection{},
	CollectionLocation{},
	Continent{},
	ContentReport{},
	Conversation{},
	Country{},
	Currency{},
	Device{},
	EngagementAnalytics{},
	Event{},
	EventAnalytics{},
	Follow{},
	GoogleType{},
	Interest{},
	InterestImage{},
	Location{},
	LocationCategory{},
	LocationClaim{},
	LocationDirectory{},
	LocationGalleryImage{},
	LocationNote{},
	LocationTag{},
	MerchantAnalytics{},
	Metric{},
	MetricSource{},
	Note{},
	NoteClosure{},
	NoteRanking{},
	NoteRepost{},
	Notification{},
	OnboardingProfile{},
	Persona{},
	Place{},
	PlaceAnalytics{},
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
	Region{},
	Relay{},
	RestrictedUsername{},
	Save{},
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
	Topic{},
	VanityPath{},
	VerificationToken{},
	Weather{},
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
