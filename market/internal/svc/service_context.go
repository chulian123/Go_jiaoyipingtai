package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"market/internal/config"
	"market/internal/database"
	"mscoin-common/msdb"
)

// ServiceContext 在这里分别注册上config ，redis，mysql的配置内容
type ServiceContext struct {
	Config      config.Config
	Cache       cache.Cache
	Db          *msdb.MsDB
	MongoClient *database.MongoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.CacheRedis, nil, cache.NewStat("coin"), nil, func(o *cache.Options) {})
	return &ServiceContext{
		Config:      c,
		Cache:       redisCache,
		Db:          database.ConnMysql(c.Mysql.DataSource),
		MongoClient: database.ConnectMongo(c.Mongo),
	}

}
