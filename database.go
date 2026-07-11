package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseModels = []interface{}{
	// Nostr related:
	Event{},
	Tag{},
	Relay{},
	// Account related:
	Account{},
	Follow{},
	Device{},
	RestrictedUsername{},
	OnboardingProfile{},
	// Calendar related:
	Calendar{},
	CalendarSubscription{},
	AccountCalendarRole{},
	CalendarCalendarEvent{},
	// Event related:
	CalendarEvent{},
	CalendarEventAnnouncement{},
	CalendarEventDraft{},
	CalendarEventCohost{},
	CalendarEventSpeaker{},
	CalendarEventNote{},
	CalendarEventRSVP{},
	CalendarEventTag{},
	CalendarEventTicketType{},
	CalendarEventCoupon{},
	CalendarEventCouponRedemption{},
	CalendarEventTicket{},
	CalendarEventTicketOrder{},
	CalendarEventTicketOrderItem{},
	CalendarEventTicketOrderPayment{},
	CalendarEventTicketOrderRefund{},
	CalendarEventTicketTypePromotion{},
	CalendarEventWallet{},
	CalendarEventWalletWithdrawal{},
	CalendarEventWalletTransaction{},
	AccountRSVPAnswers{},
	// Community related:
	Community{},
	CommunityGalleryImage{},
	CommunityMembershipTier{},
	CommunityMembershipRequest{},
	CommunityMember{},
	CommunityMembershipSubscription{},
	CommunityMembershipSubscriptionChange{},
	CommunityMembershipPayment{},
	CommunityMembershipRefund{},
	CommunityWallet{},
	CommunityWalletWithdrawal{},
	CommunityWalletTransaction{},
	CommunityNewsletter{},
	AccountCommunityRole{},
	// Collection related:
	Collection{},
	CollectionLocation{},
	CollectionSave{},
	AccountCollectionRole{},
	// Interest / personalization related:
	Interest{},
	InterestImage{},
	Cluster{},
	ClusterInterest{},
	Persona{},
	// Location related:
	Location{},
	LocationCategory{},
	LocationClaim{},
	LocationDirectory{},
	LocationGalleryImage{},
	LocationTag{},
	AccountLocationRole{},
	AccountLocationReview{},
	GoogleType{},
	// Geo related:
	Continent{},
	Country{},
	Region{},
	Place{},
	PlaceAnalytics{},
	PlaceCategoryScore{},
	PlaceClosure{},
	PlaceFoundationScore{},
	PlaceGalleryImage{},
	PlaceMetric{},
	PlacePublicChatChannel{},
	PlaceTopicScore{},
	AccountPlaceRole{},
	Brand{},
	Weather{},
	// Feed related:
	Note{},
	NoteClosure{},
	NoteRanking{},
	NoteRepost{},
	AccountNoteTracker{},
	Activity{},
	Block{},
	ContentReport{},
	LocationNote{},
	PlaceNote{},
	Reaction{},
	// Wallet related:
	AccountWallet{},
	AccountWalletTransaction{},
	AccountStripeConnect{},
	LightningContact{},
	// Metrics related:
	Category{},
	Topic{},
	Metric{},
	MetricSource{},
	// Chat related:
	ChatMembership{},
	ChatMessage{},
	Conversation{},
	PublicChatChannel{},
	PublicChatMembership{},
	PublicChatMessage{},
	// External data source related:
	SourceBots{},
	SourceCountryStats{},
	SourceFacts{},
	SourceGoogleAnalytics{},
	SourceLocationsOsm{},
	SourceLocationsAll{},
	SourceNumbeo{},
	SourceRuleoflaw{},
	SourceSpeedtest{},
	SourceSurvey{},
	// Analytics related:
	AccountAnalytics{},
	EngagementAnalytics{},
	EventAnalytics{},
	MerchantAnalytics{},
	// Others:
	Currency{},
	Notification{},
	Theme{},
	VanityPath{},
	VerificationToken{},
	WhopCompany{},
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
