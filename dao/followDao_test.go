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

func TestFollowDao_GetFollowListRedis(t *testing.T) {
	initializer.InitConfig()
	initializer.InitRedis()
	followDao := NewFollowDaoInstance()
	output, _ := followDao.GetFollowList(1)
	// （t,真实长度，期望的长度）
	assert.Equal(t, len(output), 3)
}
