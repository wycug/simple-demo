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

type FollowerDao struct {
}

// 生成实例
var (
	followerDao  *FollowerDao
	followerOnce sync.Once
)

func NewFollowerDaoInstance() *FollowerDao {
	followerOnce.Do(func() {
		followerDao = &FollowerDao{}
	})
	return followerDao
}

// 1、先根据user表 用户id
// 2、用户id 作为user_id 字段，找所有相关的follow_id数据

// 自己的粉丝列表：把自己的id作为follow_user_id传入，读取所有相关的条目

func (f FollowerDao) GetFollowerList(userID int64) ([]string, error) {
	key := fmt.Sprintf("t%v:fanslist", userID)
	return global.Rdb.SMembers(ctx, key).Result()
}

func (f FollowerDao) GetNoneFollow(userID int64) ([]string, error) {
	key1 := fmt.Sprintf("%v:fanslist", userID)
	key2 := fmt.Sprintf("%v:followlist", userID)
	return global.Rdb.SDiff(ctx, key1, key2).Result()
}

func (f FollowerDao) GetIsFollow(userID int64) ([]string, error) {
	key1 := fmt.Sprintf("%v:fanslist", userID)
	key2 := fmt.Sprintf("%v:followlist", userID)
	return global.Rdb.SInter(ctx, key1, key2).Result()
}

// 根据粉丝的id找到粉丝的信息
// 传入粉丝id列表，返回粉丝信息列表

func (f FollowerDao) GetFollowerInfoList(followRes []int64) ([]*repository.User, error) {
	// 先声明一个数组，
	var followerInfo []*repository.User
	// 传入粉丝id列表，一次查询所有粉丝信息列表，查询到的值返回到数组中
	if err := global.Db.Where("id In (?)", followRes).Find(&followerInfo).Error; err != nil {
		return nil, err
	}

	return followerInfo, nil
}
