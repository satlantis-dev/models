package models

type InterestNote struct {
	ID         uint `gorm:"primaryKey" json:"id"`
	InterestID uint `gorm:"index" json:"interestId"`
	NoteID     uint `gorm:"index" json:"noteId"`
	Note       Note `json:"note"`
}
