/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (User) TableName() string {
	return "user_info"
}
