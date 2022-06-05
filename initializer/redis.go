/**
@author: Junming ZHANG, Haojun YIN
@date: 2022/6/2
*/

package initializer

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/go-redis/redis/v8"
)

// Redis 配置Redis
func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB, // use default DB
	})
	global.Rdb = rdb
}
