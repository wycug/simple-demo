package dao

import (
	"gorm.io/gorm"
)

//存放数据定义

type UserInfo struct {
	gorm.Model
	Id            int64  `gorm:"primaryKey"`
	Name          string `gorm:"varchar(50);not null;unique"`
	Password      string `gorm:"varchar(50);not null"`
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool `json:"is_follow,omitempty"`
	Token         string
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

type CommentInfo struct {
	gorm.Model
	Id           int64 `gorm:"comment:自增主键"`
	UserID       int64
	VideoID      int64
	Content      string `gorm:"varchar(50);not null"`
	Comment_date string
}

// type CommentAPI struct {
// 	gorm.Model
// 	ID       int64
// 	User     UserAPI
// 	Content  string
// 	CreateAt string
// }

// UserAPI 主要提供给接口使用
// type UserAPI struct {
// 	gorm.Model
// 	ID            int64  `json:"id"`
// 	Name          string `json:"name"`
// 	FollowCount   int    `json:"follow_count"`
// 	FollowerCount int    `json:"follower_count"`
// 	IsFollow      bool   `json:"is_follow"`
// }

type FavoriteInfo struct {
	gorm.Model
	UserId     int64
	VideoId    int64
	IsFavorite int64
}
