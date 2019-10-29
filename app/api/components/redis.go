package components

import (
	"github.com/go-redis/redis"
	"strings"
)

var (
	Redis        *redis.Client
	RedisCluster *redis.ClusterClient
)

func InitRedis() {
	section := "redis"
	addr := Config.MustValue(section, "addr", "127.0.0.1:6379")
	password := Config.MustValue(section, "password", "")
	db := Config.MustInt(section, "password", 0)
	maxRetries := Config.MustInt(section, "max_retries", 2)
	poolSize := Config.MustInt(section, "pool_size")
	minIdleConns := Config.MustInt(section, "min_idle_conns", 5)

	if addr == "" {
		panic("缺少必要的redis配置")
	}

	opt := redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         db,
		MaxRetries: maxRetries,
	}
	if poolSize > 0 {
		opt.PoolSize = poolSize
	}
	if minIdleConns > 0 {
		opt.MinIdleConns = minIdleConns
	}

	Redis = redis.NewClient(&opt)
}

func InitRedisCluster() {
	section := "redis_cluster"
	addr := Config.MustValue(section, "addr")
	password := Config.MustValue(section, "password", "")
	maxRetries := Config.MustInt(section, "max_retries", 2)
	poolSize := Config.MustInt(section, "pool_size")
	minIdleConns := Config.MustInt(section, "min_idle_conns", 5)

	if addr == "" {
		panic("缺少必要的redis配置")
	}
	addrs := strings.Split(addr, ";")
	if len(addrs) < 1 {
		panic("redis集群地址配置错误")
	}

	opt := redis.ClusterOptions{
		Addrs: addrs,
		Password:   password,
		MaxRetries: maxRetries,
	}
	if poolSize > 0 {
		opt.PoolSize = poolSize
	}
	if minIdleConns > 0 {
		opt.MinIdleConns = minIdleConns
	}
	RedisCluster = redis.NewClusterClient(&opt)
}
