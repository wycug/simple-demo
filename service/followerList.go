/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
**/

// 传当前用户id，找dao层函数，返回列表

package service

import (
	"github.com/RaymondCode/simple-demo/util"
	"github.com/RaymondCode/simple-demo/util/myError"
	"sync"

	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/util/constant"
)

// 关注粉丝列表

type FollowerListService struct {
}

var (
	followerListService     *FollowerListService
	followerListServiceOnce sync.Once

	followDao   = dao.NewFollowDaoInstance()
	followerDao = dao.NewFollowerDaoInstance()

	noneFollow []string
	isFollow   []string

	infoList []*repository.User
)

// service层的实例 和dao没有关系
func NewFollowerDaoInstance() *FollowerListService {
	followerListServiceOnce.Do(func() {
		followerListService = &FollowerListService{}
	})
	return followerListService
}

func (f FollowerListService) FollowerList(user_id int64, opt int64) ([]*repository.User, []int64, []int64, error) {
	// 先读取操作数（读关注列表还是读粉丝列表）
	//  读关注列表，通过用户id，调dao层查询关注列表函数
	lock.Lock()
	defer lock.Unlock()

	if opt == constant.FOLLOWLIST {
		followNum, err := followDao.GetFollowList(user_id)
		if err != nil {
			return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), myError.NewError(constant.GetFollowIDListError, constant.Msg(constant.GetFollowIDListError))
		}
		infoList, err = followDao.GetFollowInfoList(util.Str2Int64(followNum))
		if err != nil {
			return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), myError.NewError(constant.GetFollowListError, constant.Msg(constant.GetFollowListError))
		}
	} else if opt == constant.FANSLIST {
		//  读粉丝列表，通过用户id，调dao层查询分数列表函数,返回自己的粉丝列表用户信息
		followerNum, err := followerDao.GetFollowerList(user_id)
		if err != nil {
			return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), myError.NewError(constant.GetFollowerIDListError, constant.Msg(constant.GetFollowerIDListError))
		}
		infoList, err = followerDao.GetFollowerInfoList(util.Str2Int64(followerNum))
		if err != nil {
			return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), myError.NewError(constant.GetFollowerListError, constant.Msg(constant.GetFollowerListError))
		}

		noneFollow, err = followerDao.GetNoneFollow(user_id)
		if err != nil {
			return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), myError.NewError(constant.GetNoneFollowListError, constant.Msg(constant.GetNoneFollowListError))
		}
		isFollow, err = followerDao.GetIsFollow(user_id)
		if err != nil {
			return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), myError.NewError(constant.GetIsFollowListError, constant.Msg(constant.GetIsFollowListError))
		}
	} else {
		return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), myError.NewError(constant.OptParameterError, constant.Msg(constant.OptParameterError))
	}
	return infoList, util.Str2Int64(noneFollow), util.Str2Int64(isFollow), nil
}
