package models

type Relay struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	AccountID uint       `gorm:"index" json:"accountId"`
	Account   AccountDTO `json:"account"`
	EventID   *uint      `gorm:"index" json:"eventId"`
	Event     Event      `json:"event"`
	Address   string     `json:"address"`
	Read      bool       `json:"read"`
	Write     bool       `json:"write"`
}
