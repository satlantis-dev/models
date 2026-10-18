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
	CalendarEventType{},
	CalendarEvent{},
	CalendarEventAnnouncement{},
	CalendarEventDraft{},
	CalendarEventCohost{},
	CalendarEventSpeaker{},
	CalendarEventNote{},
	CalendarEventRSVP{},
	CalendarEventTag{},
	CalendarEventTicketType{},
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
	CommunityMemberAuditLog{},
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
	LocationLocationTag{},
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
	// Plan related:
	Plan{},
	PlanMember{},
	PlanSubscription{},
	PlanSubscriptionChange{},
	PlanPayment{},
	// Coupon related:
	Coupon{},
	CouponRedemption{},
	// Others:
	Banner{},
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

	if err := SetupJoinTables(db); err != nil {
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

// SetupJoinTables registers the join models of many2many relationships whose join table has
// more than GORM's default (two key columns): extra columns or indexes. Without it GORM treats
// the table as a plain join table and AutoMigrate skips the extras. Call it on any *gorm.DB that
// isn't opened with SetupDatabase (tests, scripts) before AutoMigrate.
func SetupJoinTables(db *gorm.DB) error {
	joins := []struct {
		model     interface{}
		field     string
		joinModel interface{}
	}{
		// Extra columns: featured, members_only, tier_ids_only.
		{&Calendar{}, "Events", &CalendarCalendarEvent{}},
		// Index on location_id.
		{&Location{}, "LocationTags", &LocationLocationTag{}},
		{&LocationTag{}, "Locations", &LocationLocationTag{}},
	}
	for _, j := range joins {
		if err := db.SetupJoinTable(j.model, j.field, j.joinModel); err != nil {
			return err
		}
	}
	return nil
}
