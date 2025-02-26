package models

import "time"

type Follow struct {
	FollowerID  uint      `gorm:"index:idx_follower_following,unique"`
	FollowingID uint      `gorm:"index:idx_follower_following,unique"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}
