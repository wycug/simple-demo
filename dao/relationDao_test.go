/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package dao

import (
	"github.com/RaymondCode/simple-demo/initializer"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRelationDao_AddFollowRelation(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	relationDao := NewRelationDaoInstance()
	output, _ := relationDao.AddFollowRelation(51, 5000)
	expectOutput := "success"
	assert.Equal(t, output, expectOutput)
}

func TestRelationDao_RemoveFollowRelation(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	relationDao := NewRelationDaoInstance()
	output, _ := relationDao.RemoveFollowRelation(50, 5000)
	expectOutput := "success"
	assert.Equal(t, output, expectOutput)
}

func TestRelationDao_SearchFollowRelation(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	relationDao := NewRelationDaoInstance()
	output, _ := relationDao.SearchFollowRelation(50, 5000)
	exceptOutput := true
	assert.Equal(t, output, exceptOutput)
}
