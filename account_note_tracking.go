package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type AccountNoteTracking struct {
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
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported type: %T", value)
	}
	return json.Unmarshal(bytes, fn)
}

func (fn FeedNotes) Value() (driver.Value, error) {
	return json.Marshal(fn)
}
