/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
**/

// 传当前用户id，找dao层函数，返回列表

package service

import (
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

	noneFollow []int64
	isFollow   []int64

	infoList []*repository.User
)

// service层的实例 和dao没有关系
func NewFollowerDaoInstance() *FollowerListService {
	followerListServiceOnce.Do(func() {
		followerListService = &FollowerListService{}
	})
	return followerListService
}

// 读取关注or粉丝列表操作,传入当前(用户id,操作数),返回列表
func (f FollowerListService) FollowerList(user_id int64, opt int64) ([]*repository.User, []int64, []int64, error) {
	// 先读取操作数（读关注列表还是读粉丝列表）
	//  读关注列表，通过用户id，调dao层查询关注列表函数
	if opt == constant.FOLLOWLIST {

		followNum, _ := followDao.GetFollowList(user_id)
		infoList, _ = followDao.GetFollowInfoList(followNum)
	} else if opt == constant.FANSLIST {
		//  读粉丝列表，通过用户id，调dao层查询分数列表函数,返回自己的粉丝列表用户信息
		followerNum, _ := followerDao.GetFollowerList(user_id)
		infoList, _ = followerDao.GetFollowerInfoList(followerNum)

		// 读自己的关注列表
		followNum, _ := followDao.GetFollowList(user_id)

		// 粉丝中自己没关注的有哪些
		noneFollow = NoneFollow(followNum, followerNum)

		// 粉丝中自己已经关注的有哪些
		isFollow = IsFollow(followerNum, followNum)

	} else {
		return infoList, noneFollow, isFollow, myError.NewError(constant.OptParameterError, constant.Msg(constant.OptParameterError))
	}
	return infoList, noneFollow, isFollow, nil
}

// 数组求差集操作,自己没关注自己的粉丝
func NoneFollow(A, B []int64) []int64 {
	if len(A) < 1 || len(B) < 1 {
		return A
	}
	result := make([]int64, 0)
	// 去重
	flagMap := make(map[int64]bool, 0)
	for _, a := range A {
		if _, ok := flagMap[a]; ok {
			continue
		}
		flagMap[a] = true
		flag := true
		for _, b := range B {
			if b == a {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, a)
		}
	}
	return result
}

// 求交集
func IsFollow(A, B []int64) []int64 {
	if len(A) < 1 || len(B) < 1 {
		return []int64{}
	}
	result := make([]int64, 0)
	// 去重
	flagMap := make(map[int64]bool, 0)
	for _, a := range A {
		if _, ok := flagMap[a]; ok {
			continue
		}
		flagMap[a] = true
		for _, b := range B {
			if b == a {
				result = append(result, a)
				break
			}
		}
	}
	return result
}
