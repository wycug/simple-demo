/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package service

import (
	"github.com/RaymondCode/simple-demo/initializer"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRelationService_Follow(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	relationService := NewRelationServiceInstance()
	output, _ := relationService.Follow(16, 17)
	exceptOutput := "success"
	assert.Equal(t, output, exceptOutput)
}

func TestRelationService_Follow2(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	relationService := NewRelationServiceInstance()
	output, _ := relationService.Follow(16, 17)
	exceptOutput := "Already followed"
	assert.Equal(t, output, exceptOutput)
}

func TestRelationService_UnFollow(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	relationService := NewRelationServiceInstance()
	output, _ := relationService.UnFollow(16, 17)
	exceptOutput := "success"
	assert.Equal(t, output, exceptOutput)
}

func TestRelationService_UnFollow2(t *testing.T) {
	initializer.InitConfig()
	initializer.InitDataBase()
	relationService := NewRelationServiceInstance()
	output, _ := relationService.UnFollow(16, 17)
	exceptOutput := "Not follow yet"
	assert.Equal(t, output, exceptOutput)
}
