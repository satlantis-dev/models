package models

import "time"

type Follow struct {
	FollowerID  uint      `gorm:"not null;uniqueIndex:idx_follower_following;index"`
	FollowingID uint      `gorm:"not null;uniqueIndex:idx_follower_following;index"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}
