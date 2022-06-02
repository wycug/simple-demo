/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package repository

import "gorm.io/gorm"

type FollowRelation struct {
	gorm.Model
	ID           int64 `gorm:"column:id"`
	FollowFromID int64 `gorm:"column:user_id"`
	FollowToID   int64 `gorm:"column:follow_user_id"`
}

func (FollowRelation) TableName() string {
	return "follow_info"
}
