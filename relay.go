package models

type Relay struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	AccountID uint       `gorm:"index" json:"accountId"`
	Account   AccountDTO `json:"account"`
	EventID   *uint      `gorm:"index" json:"eventId"`
	Address   string     `gorm:"index" json:"address"`
	Read      bool       `json:"read"`
	Write     bool       `json:"write"`
	// AccountCount is the number of distinct accounts that currently list this Address as one
	// of their relays. Denormalized (duplicated across every row sharing this Address) and
	// recomputed after writes rather than incremented/decremented in place.
	AccountCount int `gorm:"default:0" json:"accountCount"`
}
