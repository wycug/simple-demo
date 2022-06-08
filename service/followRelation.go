/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/util/constant"
	"github.com/RaymondCode/simple-demo/util/myError"
	"sync"
)

type RelationService struct {
}

var (
	relationService     *RelationService
	relationServiceOnce sync.Once

	userDao          = dao.NewUserDaoInstance()
	relationRedisDao = dao.NewRelationRedisDaoInstance()

	lock sync.Mutex
)

func NewRelationServiceInstance() *RelationService {
	relationServiceOnce.Do(func() {
		relationService = &RelationService{}
	})
	return relationService
}

func (s RelationService) Follow(followFromID int64, followToID int64) (string, error) {

	lock.Lock()
	defer lock.Unlock()

	flag, err := relationRedisDao.SearchFollowRelation(followFromID, followToID)
	if err != nil {
		return constant.Msg(constant.FollowFailed), err
	}

	if !flag {
		tx := global.Db.Begin()
		if _, err := relationRedisDao.AddFollowRelation(followFromID, followToID); err != nil {
			tx.Rollback()
			return err.Error(), err
		}
		// 修改当前用户的关注数 ++
		if err := userDao.FollowNumChange(followFromID, constant.INCREASE); err != nil {
			tx.Rollback()
			return err.Error(), err
		}
		// 修改被关注的用户的粉丝数 ++
		if err := userDao.FansNumChange(followToID, constant.INCREASE); err != nil {
			tx.Rollback()
			return err.Error(), err
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return constant.Msg(constant.TransactionCommitError), myError.NewError(constant.TransactionCommitError, constant.Msg(constant.TransactionCommitError))
		}
	}

	//返回
	return constant.Msg(constant.Success), nil
}

func (s RelationService) UnFollow(followFromID int64, followToID int64) (string, error) {

	lock.Lock()
	defer lock.Unlock()

	flag, err := relationRedisDao.SearchFollowRelation(followFromID, followToID)
	if err != nil {
		return constant.Msg(constant.UnfollowFailed), err
	}

	if flag {
		tx := global.Db.Begin()
		// 如果已经关注，则继续处理
		// 修改当前用户的关注数 --
		if err := userDao.FollowNumChange(followFromID, constant.DECREASE); err != nil {
			tx.Rollback()
			return err.Error(), err
		}
		// 修改被关注用户的粉丝数 --
		if err := userDao.FansNumChange(followToID, constant.DECREASE); err != nil {
			tx.Rollback()
			return err.Error(), err
		}
		// 删除数据库中的关系
		if _, err := relationRedisDao.RemoveFollowRelation(followFromID, followToID); err != nil {
			tx.Rollback()
			return err.Error(), err
		}

		if err := tx.Commit().Error; err != nil {
			tx.Rollback()
			return constant.Msg(constant.TransactionCommitError), myError.NewError(constant.TransactionCommitError, constant.Msg(constant.TransactionCommitError))
		}
	}
	// 返回
	return constant.Msg(constant.Success), nil
}
