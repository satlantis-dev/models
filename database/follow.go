package database

type Follow struct {
	FollowerID  uint `gorm:"index:idx_follower_following,unique"`
	FollowingID uint `gorm:"index:idx_follower_following,unique"`
}
