package models

type Relay struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	AccountID uint       `gorm:"index" json:"accountId"`
	Account   AccountDTO `gorm:"-:migration" json:"account"`
	EventID   *uint      `gorm:"index" json:"eventId"`
	Event     *Event     `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE" json:"-"`
	Address   string     `gorm:"index" json:"address"`
	Read      bool       `json:"read"`
	Write     bool       `json:"write"`
}
