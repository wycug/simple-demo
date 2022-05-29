/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/28
**/

package service

import (
	"testing"

	"github.com/RaymondCode/simple-demo/initializer"
	"github.com/RaymondCode/simple-demo/util/constant"
	"github.com/magiconair/properties/assert"
)

// 测试自己的粉丝列表
func TestFollowerList_FollowerList(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	followerListService := NewFollowerDaoInstance()
	output, _, _, _ := followerListService.FollowerList(20, constant.FANSLIST)
	assert.Equal(t, len(output), 2)

}

// 测试关注列表
func TestFollowerList_FollowerList2(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	followerListService := NewFollowerDaoInstance()
	output, _, _, _ := followerListService.FollowerList(17, constant.FOLLOWLIST)
	assert.Equal(t, len(output), 3)
}
