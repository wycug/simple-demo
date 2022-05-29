/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package dao

import (
	"github.com/RaymondCode/simple-demo/initializer"
	"github.com/RaymondCode/simple-demo/util/constant"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestUserDao_FollowNumChange(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	userDao := NewUserDaoInstance()
	output := userDao.FollowNumChange(17, constant.DECREASE)
	assert.Equal(t, output, nil)
}

func TestUserDao_FansNumChange(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	userDao := NewUserDaoInstance()
	output := userDao.FansNumChange(17, constant.INCREASE)
	assert.Equal(t, output, nil)
}
