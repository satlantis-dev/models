package models

type AccountLocationReview struct {
	ID           uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID    uint         `gorm:"uniqueIndex:idx_account_location_collection" json:"accountId"`
	Account      AccountDTO   `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	LocationID   uint         `gorm:"uniqueIndex:idx_account_location_collection" json:"locationId"`
	Location     *LocationDTO `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	CollectionID *uint        `gorm:"uniqueIndex:idx_account_location_collection" json:"collectionId"`
	Collection   *Collection  `gorm:"foreignKey:CollectionID" json:"collection,omitempty"`
	ReviewText   string       `json:"reviewText"`
	IsPositive   bool         `gorm:"not null;default:true" json:"isPositive"`
}

func (AccountLocationReview) TableName() string {
	return "account_location_reviews"
}
