/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/repository"
	"gorm.io/gorm"
	"sync"
)

type RelationDao struct {
}

var (
	relationDao  *RelationDao
	relationOnce sync.Once
)

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(func() {
		relationDao = &RelationDao{}
	})
	return relationDao
}

//插入关注关系

func (d RelationDao) AddFollowRelation(followFromId int64, followToId int64) (string, error) {
	// 封装关注关系
	followRelation := &repository.FollowRelation{
		FollowFromID: followFromId,
		FollowToID:   followToId,
	}
	//插入关注关系
	global.Db.Select("FollowFromID", "FollowToID").Create(followRelation)

	return "success", nil
}

// 删除关注关系

func (d RelationDao) RemoveFollowRelation(followFromId int64, followToId int64) (string, error) {
	// 封装关注关系
	global.Db.Where("user_id = ? and follow_user_id = ?",
		followFromId, followToId).Delete(&repository.FollowRelation{})

	return "success", nil
}

// 获取关注关系

func (d RelationDao) SearchFollowRelation(followFromId int64, followToId int64) (bool, error) {
	// 查询关系中是否有 followFromId -> followToId, 返回满足followFromId -> followToId的数据集
	e := global.Db.Where("user_id = ? and follow_user_id = ?",
		followFromId, followToId).First(&repository.FollowRelation{}).Error
	if e == gorm.ErrRecordNotFound {
		fmt.Println("查询失败")
		return false, nil
	}
	return true, nil
}
