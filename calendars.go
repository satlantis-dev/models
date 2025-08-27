package models

type Calendar struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Banner      string          `json:"banner"`
	AccountID   uint            `json:"account_id"`
	Account     *AccountDTO     `gorm:"foreignKey:AccountID" json:"account,omitempty"`
	Events      []CalendarEvent `gorm:"many2many:calendar_events;" json:"events"`
}
