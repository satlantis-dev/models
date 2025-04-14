package models

type LocationDirectory struct {
	Code        string  `gorm:"index;primaryKey" json:"code"`
	Name        string  `json:"name"`
	Npub        string  `json:"npub"`
	AccountID   uint    `gorm:"index" json:"accountId"`
	Account     Account `gorm:"foreignKey:AccountID" json:"account"`
	Description string  `json:"description"`
}
