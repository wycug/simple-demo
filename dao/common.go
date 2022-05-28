package dao

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Id            int64  `gorm:"primaryKey"`
	Name          string `gorm:"varchar(50);not null;unique"`
	Password      string `gorm:"varchar(50);not null"`
	FollowCount   int64
	FollowerCount int64
	Token         string `gorm:"not null;index;"`
}

type FollowInfo struct {
	gorm.Model
	UserId       int64 `gorm:"not null;index;"`
	FollowUserId int64 `gorm:"not null;index;"`
}

type VideoInfo struct {
	gorm.Model
	Id            int64 `gorm:"primaryKey"`
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	Title         string
}

type FavoriteInfo struct {
	gorm.Model
	UserId     int64
	VideoId    int64
	IsFavorite int64
}
