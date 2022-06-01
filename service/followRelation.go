/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/util/constant"
	"github.com/RaymondCode/simple-demo/util/myError"
	"sync"
)

type RelationService struct {
}

var (
	relationService     *RelationService
	relationServiceOnce sync.Once

	userDao     = dao.NewUserDaoInstance()
	relationDao = dao.NewRelationDaoInstance()
)

func NewRelationServiceInstance() *RelationService {
	relationServiceOnce.Do(func() {
		relationService = &RelationService{}
	})
	return relationService
}

func (s RelationService) Follow(followFromID int64, followToID int64) (string, error) {
	// 判断是否已经关注，如果关注过，则直接返回
	flag, _ := relationDao.SearchFollowRelation(followFromID, followToID)
	if flag {
		return "Already followed", myError.NewError(constant.AlreadyFollowedError, constant.Msg(constant.AlreadyFollowedError))
	}

	// 如果没有关注，则继续处理
	// 先把关系插入到数据库
	if _, err := relationDao.AddFollowRelation(followFromID, followToID); err != nil {
		return err.Error(), err
	}
	// 修改当前用户的关注数 ++
	if err := userDao.FollowNumChange(followFromID, constant.INCREASE); err != nil {
		return err.Error(), err
	}
	// 修改被关注的用户的粉丝数 ++
	if err := userDao.FansNumChange(followToID, constant.INCREASE); err != nil {
		return err.Error(), err
	}

	//返回
	return "success", nil
}

func (s RelationService) UnFollow(followFromID int64, followToID int64) (string, error) {
	// 判断是否已经关注，如果没关注过，则直接返回
	flag, _ := relationDao.SearchFollowRelation(followFromID, followToID)
	if !flag {
		return "Not follow yet", myError.NewError(constant.NotFollowYetError, constant.Msg(constant.NotFollowYetError))
	}

	// 如果已经关注，则继续处理
	// 修改当前用户的关注数 --
	if err := userDao.FollowNumChange(followFromID, constant.DECREASE); err != nil {
		return err.Error(), err
	}
	// 修改被关注用户的粉丝数 --
	if err := userDao.FansNumChange(followToID, constant.DECREASE); err != nil {
		return err.Error(), err
	}
	// 删除数据库中的关系
	if _, err := relationDao.RemoveFollowRelation(followFromID, followToID); err != nil {
		return err.Error(), err
	}

	// 返回
	return "success", nil
}
