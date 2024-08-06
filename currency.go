package models

type Currency struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Code         string  `gorm:"uniqueIndex" json:"code"`
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	ExchangeRate float64 `json:"exchangeRate"`
}
