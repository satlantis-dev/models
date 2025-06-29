package models

type AccountLocationReview struct {
	AccountID  uint         `gorm:"primaryKey;index" json:"accountId"`
	Account    AccountDTO   `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	LocationID uint         `gorm:"primaryKey;index" json:"locationId"`
	Location   *LocationDTO `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	ReviewText string       `json:"reviewText"`
	IsPositive bool         `gorm:"not null;default:true" json:"isPositive"`
}

func (AccountLocationReview) TableName() string {
	return "account_location_reviews"
}
