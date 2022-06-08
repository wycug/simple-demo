/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/5/27
**/

package global

import (
	"github.com/RaymondCode/simple-demo/config"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)

var (
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
)
