package models

type AccountNoteTracker struct {
	AccountID        uint       `gorm:"index;primaryKey" json:"accountId"`
	Account          AccountDTO `gorm:"foreignKey:AccountID" json:"account"`
	BloomFilter      []byte     `json:"bloomFilter"`
	FeedNotes        string     `gorm:"type:jsonb" json:"feedNotes"`
	BloomFilterCount int        `gorm:"default:0" json:"bloomFilterCount"`
}
