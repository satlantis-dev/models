package database

type AccountLocationRating struct {
	AccountID  uint   `gorm:"primaryKey" json:"accountId"`
	LocationID uint   `gorm:"primaryKey" json:"locationId"`
	Review     string `json:"review"`
	Ratings    string `gorm:"type:jsonb" json:"ratings"`
}
