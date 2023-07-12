package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql      MysqlConfig
	CacheRedis cache.CacheConf
	Captcha    CaptchaConf
	JWT        AutoConfig
	MarketRpc  zrpc.RpcClientConf
}

type AutoConfig struct {
	AccessSecret string
	AccessExpire int64
}

type CaptchaConf struct {
	Vid string
	Key string
}

type MysqlConfig struct {
	DataSource string
}
