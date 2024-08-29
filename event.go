package models

type Event struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	NostrID    string `gorm:"column:nostr_id;index" json:"nostrId"`
	CreatedAt  int64  `json:"createdAt"`
	Content    string `gorm:"type:text" json:"content"`
	Kind       uint   `gorm:"index" json:"kind"`
	PubKey     string `gorm:"type:text" json:"pubkey"`
	Sig        string `gorm:"type:text" json:"sig"`
	Tags       []Tag  `gorm:"foreignKey:EventID" json:"tags"`
	Reconciled bool   `gorm:"default:false" json:"reconciled"`
}
