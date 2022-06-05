/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/6/2
*/

package dao

import (
	"github.com/RaymondCode/simple-demo/initializer"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRelationRedisDao_AddFollowRelation(t *testing.T) {
	initializer.InitConfig()
	initializer.InitRedis()
	relationRedisDao := NewRelationRedisDaoInstance()
	output, _ := relationRedisDao.AddFollowRelation(1, 9)
	expectOutput := "success"
	assert.Equal(t, output, expectOutput)
}

func TestRelationRedisDao_RemoveFollowRelation(t *testing.T) {
	initializer.InitConfig()
	initializer.InitRedis()
	relationRedisDao := NewRelationRedisDaoInstance()
	output, _ := relationRedisDao.RemoveFollowRelation(1, 9)
	expectOutput := "success"
	assert.Equal(t, output, expectOutput)
}

func TestRelationRedisDao_SearchFollowRelation(t *testing.T) {
	initializer.InitConfig()
	initializer.InitRedis()
	relationRedisDao := NewRelationRedisDaoInstance()
	output, _ := relationRedisDao.SearchFollowRelation(1, 9)
	assert.Equal(t, output, true)
}
