package models

type Cluster struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Label     string `gorm:"type:text" json:"label"`
	UserCount int    `gorm:"default:0" json:"userCount"`
	Active    bool   `gorm:"default:true" json:"active"`
}

type ClusterInterest struct {
	ClusterID  uint    `gorm:"primaryKey" json:"clusterId"`
	InterestID uint    `gorm:"primaryKey" json:"interestId"`
	Score      float64 `json:"score"`
}
