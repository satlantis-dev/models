package models

type AccountNoteTracker struct {
	AccountID        uint     `gorm:"primaryKey;autoIncrement:false" json:"accountId"`
	Account          *Account `gorm:"foreignKey:AccountID;constraint:OnDelete:CASCADE;" json:"account"`
	BloomFilter      []byte   `json:"bloomFilter"`
	FeedNotes        string   `gorm:"type:jsonb" json:"feedNotes"`
	BloomFilterCount int      `gorm:"default:0" json:"bloomFilterCount"`
}
