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

func TestFollowDao_FollowList(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	followDao := NewFollowDaoInstance()
	output, _ := followDao.GetFollowList(17)
	// （t,真实长度，期望的长度）
	assert.Equal(t, len(output), 3)
}
