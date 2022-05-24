package dao

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Id             uint64 `gorm:"primaryKey"`
	Name           string `gorm:"varchar(50);not null;unique"`
	Password       string `gorm:"varchar(50);not null"`
	FollowCount    uint64
	Follower_count uint64
}

type FollowInfo struct {
	gorm.Model
	UserId       uint64 `gorm:"not null;index;"`
	FollowUserId uint64 `gorm:"not null;index;"`
}
