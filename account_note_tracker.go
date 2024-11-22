package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type AccountNoteTracker struct {
	AccountID   uint       `gorm:"index;primaryKey" json:"accountId"`
	Account     AccountDTO `gorm:"foreignKey:AccountID" json:"account"`
	BloomFilter []byte     `json:"bloomFilter"`
	FeedNotes   []FeedNote `gorm:"type:jsonb" json:"feedNotes"`
}

type FeedNotes []FeedNote

func (fn *FeedNotes) Scan(value interface{}) error {
	if value == nil {
		*fn = FeedNotes{}
		return nil
	}
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, fn)
	case string:
		return json.Unmarshal([]byte(v), fn)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

func (fn FeedNotes) Value() (driver.Value, error) {
	return json.Marshal(fn)
}
