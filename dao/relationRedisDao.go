/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/6/2
*/

package dao

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/util/constant"
	"github.com/RaymondCode/simple-demo/util/myError"
	"sync"
)

type RelationRedisDao struct {
}

var (
	relationRedisDao  *RelationRedisDao
	relationRedisOnce sync.Once

	ctx = context.Background()
)

func NewRelationRedisDaoInstance() *RelationRedisDao {
	relationRedisOnce.Do(func() {
		relationRedisDao = &RelationRedisDao{}
	})
	return relationRedisDao
}

//插入关注关系

func (d RelationRedisDao) AddFollowRelation(followFromId int64, followToId int64) (string, error) {
	// 构建关注列表的键和粉丝列表的键
	pipe := global.Rdb.TxPipeline()

	// followFromId 的 关注列表中添加 followToId
	key1 := fmt.Sprintf("%v:followlist", followFromId)
	if err := pipe.SAdd(ctx, key1, followToId).Err(); err != nil {
		return "follow failed", err
	}

	// followToId 的 粉丝列表中添加 followFromId
	key2 := fmt.Sprintf("%v:fanslist", followToId)
	if err := pipe.SAdd(ctx, key2, followFromId).Err(); err != nil {
		return "follow failed", myError.NewError(constant.FollowFailed, constant.Msg(constant.FollowFailed))
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return "redis transaction error", err
	}
	return "success", nil
}

// 删除关注关系

func (d RelationRedisDao) RemoveFollowRelation(followFromId int64, followToId int64) (string, error) {
	// 构建关注列表的键和粉丝列表的键
	pipe := global.Rdb.TxPipeline()
	// followFromId 的 关注列表中删除 followToId
	key1 := fmt.Sprintf("%v:followlist", followFromId)
	if err := pipe.SRem(ctx, key1, followToId).Err(); err != nil {
		return "unfollow failed", myError.NewError(constant.UnfollowFailed, constant.Msg(constant.UnfollowFailed))
	}

	// followToId 的 粉丝列表中删除 followFromId
	key2 := fmt.Sprintf("%v:fanslist", followToId)
	if err := pipe.SRem(ctx, key2, followFromId).Err(); err != nil {
		return "unfollow failed", myError.NewError(constant.UnfollowFailed, constant.Msg(constant.UnfollowFailed))
	}

	if _, err := pipe.Exec(ctx); err != nil {
		return "redis transaction error", err
	}

	return "success", nil
}

// 查询关系中是否有 followFromId -> followToId, 返回满足followFromId -> followToId的数据集

func (d RelationRedisDao) SearchFollowRelation(followFromId int64, followToId int64) (bool, error) {
	// 查询关系中是否有 followFromId -> followToId, 返回满足followFromId -> followToId的数据集
	key := fmt.Sprintf("%v:followlist", followFromId)
	return global.Rdb.SIsMember(ctx, key, followToId).Result()
}
