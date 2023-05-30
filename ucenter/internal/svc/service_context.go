package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"ucenter/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.CacheRedis, nil, cache.NewStat("coin"), nil, func(o *cache.Options) {})
	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
	}

}
