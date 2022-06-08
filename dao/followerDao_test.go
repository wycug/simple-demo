/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
**/

package dao

import (
	"github.com/RaymondCode/simple-demo/util"
	"testing"

	"github.com/RaymondCode/simple-demo/initializer"
	"github.com/magiconair/properties/assert"
)

func TestFollowerDao_GetFollowerList(t *testing.T) {
	initializer.InitConfig()
	initializer.InitRedis()
	followerDao := NewFollowerDaoInstance()
	output, _ := followerDao.GetFollowerList(5)
	// （t,真实长度，期望的长度）
	assert.Equal(t, len(output), 3)
}

func TestFollowerDao_FollowerInfoList(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	followerDao := NewFollowerDaoInstance()

	// 找18的粉丝列表
	input, _ := followerDao.GetFollowerList(18)
	// 输入粉丝列表，输出对应粉丝信息
	output, _ := followerDao.GetFollowerInfoList(util.Str2Int64(input))
	assert.Equal(t, len(output), len(input))
}


