/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
**/

package dao

import (
	"fmt"
	"sync"

	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/repository"
)

type FollowDao struct {
}

// 生成实例
var (
	followDao  *FollowDao
	followOnce sync.Once
)

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(func() {
		followDao = &FollowDao{}
	})
	return followDao
}

// 1、先根据user表 用户id
// 2、用户id 作为user_id 字段，找所有相关的follow_id数据
// 关注列表

func (f FollowDao) GetFollowList(userID int64) ([]string, error) {
	key := fmt.Sprintf("%v:followlist", userID)
	return global.Rdb.SMembers(ctx, key).Result()
}

// 传入关注id列表，返回关注信息列表

func (f FollowDao) GetFollowInfoList(followRes []int64) ([]*repository.User, error) {
	// 先声明一个数组，
	var followerInfo []*repository.User
	// 传入关注id列表，一次查询所有关注信息列表，查询到的值返回到数组中
	if err := global.Db.Where("id In (?)", followRes).Find(&followerInfo).Error; err != nil {
		return nil, err
	}

	return followerInfo, nil
}
