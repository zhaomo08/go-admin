// 缓存redis配置初始化
// author xiaoRui

package redis

import (
	"admin-api/common/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

// 初始化连接
func SetupRedisDb() error {
	//var ctx = context.Background()
	//RedisDb = redis.NewClient(&redis.Options{Addr: config.Config.Redis.Address, Password: config.Config.Redis.Password, DB: 0})
	//_, err := RedisDb.Ping(ctx).Result()
	//if err != nil {
	//	return err
	//}
	//return nil
	var ctx = context.Background()

	fmt.Printf("Connecting to Redis at %s with password %s\n", config.Config.Redis.Address, config.Config.Redis.Password)

	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       0,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return err
	}
	return nil
}
