package cache

import (
	"fmt"
	"sync"
	"github.com/go-redis/redis"
	"toy/marmot/web/query-log-echo/launch/config"
	"time"
)

type CommonRedisClient struct {
	*redis.Client
}

var (
	CommonRedis    *CommonRedisClient
	redisClientOnce sync.Once
)

func InitRedisPool(){
	redisClientOnce.Do(func() {
		host := config.Gcfg.GetString("redis_common.host")
		port := config.Gcfg.GetInt("redis_common.port")
		addr := fmt.Sprintf("%s:%d", host, port)
		fmt.Printf("addr:%s", addr)
		connectTimeout := config.Gcfg.GetInt("redis_common.connect_timeout")
		readTimeout := config.Gcfg.GetInt("redis_common.read_timeout")
		writeTimeout := config.Gcfg.GetInt("redis_common.write_timeout")
		maxActive := config.Gcfg.GetInt("redis_common.max_active")
		idleTimeout := config.Gcfg.GetInt("redis_common.idle_timeout")

		CommonRedis = &CommonRedisClient{Client:redis.NewClient(&redis.Options{
			Addr:     addr,
			DialTimeout: time.Duration(connectTimeout )* time.Second,
			ReadTimeout:time.Duration(readTimeout)*time.Millisecond,
			WriteTimeout: time.Duration(writeTimeout)*time.Millisecond,
			IdleTimeout: time.Duration(idleTimeout)*time.Millisecond,
			PoolSize: maxActive})}
	})
}