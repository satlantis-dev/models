package database

type AccountPlaceRating struct {
	AccountID uint   `gorm:"primaryKey" json:"accountId"`
	PlaceID   uint   `gorm:"primaryKey" json:"placeId"`
	Review    string `json:"review"`
	Ratings   string `gorm:"type:jsonb" json:"ratings"`
}
