/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package dao

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/util/constant"
	"github.com/RaymondCode/simple-demo/util/myError"
	"sync"
)

type UserDao struct {
}

var (
	userDao  *UserDao
	userOnce sync.Once
)

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (d UserDao) FansNumChange(userId int64, opt int64) error {
	//查当前的粉丝数
	followerCount := GetFansNumByID(userId)
	// ++ / --
	if opt == constant.INCREASE {
		followerCount++
	} else if opt == constant.DECREASE {
		followerCount--
	} else {
		return myError.NewError(constant.OptParameterError, constant.Msg(constant.OptParameterError))
	}

	//写数据库
	if err := global.Db.Model(&repository.User{ID: userId}).Update("follower_count", followerCount).Error; err != nil {
		return err
	}

	return nil
}

func GetFansNumByID(userId int64) int64 {
	user := &repository.User{
		ID: userId,
	}
	global.Db.Where("id = ?", userId).Debug().First(&user)

	return user.FollowerCount
}

func (d UserDao) FollowNumChange(userId int64, opt int64) error {
	//查当前的关注数
	followCount := GetFollowNumByID(userId)
	//++ / --
	if opt == constant.INCREASE {
		followCount++
	} else if opt == constant.DECREASE {
		followCount--
	} else {
		return myError.NewError(constant.OptParameterError, constant.Msg(constant.OptParameterError))
	}
	//写数据库
	if err := global.Db.Model(&repository.User{ID: userId}).Update("follow_count", followCount).Error; err != nil {
		return err
	}

	return nil
}

func GetFollowNumByID(userId int64) int64 {
	user := &repository.User{
		ID: userId,
	}
	global.Db.Where("id = ?", userId).Debug().First(&user)

	return user.FollowCount
}
