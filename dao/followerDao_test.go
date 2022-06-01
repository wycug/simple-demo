/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
**/

package dao

import (
	"testing"

	"github.com/RaymondCode/simple-demo/initializer"
	"github.com/magiconair/properties/assert"
)

func TestFollwerDao_FollowerList(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	followerDao := NewFollowerDaoInstance()
	output, _ := followerDao.GetFollowerList(18)
	assert.Equal(t, len(output), 4)
}

func TestFollowerDao_FollowerInfoList(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	followerDao := NewFollowerDaoInstance()

	// 找18的粉丝列表
	input, _ := followerDao.GetFollowerList(18)
	// 输入粉丝列表，输出对应粉丝信息
	output, _ := followerDao.GetFollowerInfoList(input)
	assert.Equal(t, len(output), len(input))
}
