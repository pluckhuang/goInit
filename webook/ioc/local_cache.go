package ioc

import (
	"gitee.com/geekbang/basic-go/webook/config"
	"github.com/patrickmn/go-cache"
)

func InitLocalCache() *Cache {
	localCacheClient := cache.New(config.Config.LocalCache.DefaultExpiration, config.Config.LocalCache.CleanupInterval)
	return localCacheClient
}
