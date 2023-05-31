package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"mscoin-common/msdb"
	"ucenter/database"
	"ucenter/internal/config"
)

// ServiceContext 在这里分别注册上config ，redis，mysql的配置内容
type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *msdb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.CacheRedis, nil, cache.NewStat("coin"), nil, func(o *cache.Options) {})
	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		Db:     database.ConnMysql(c.Mysql.DataSource),
	}

}
