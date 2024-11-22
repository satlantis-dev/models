package models

type AccountNoteTracking struct {
	AccountID   uint       `gorm:"index;primaryKey" json:"accountId"`
	Account     AccountDTO `gorm:"foreignKey:AccountID" json:"account"`
	BloomFilter []byte     `json:"bloomFilter"`
	FeedNotes   []FeedNote `gorm:"type:jsonb" json:"feedNotes"`
}
