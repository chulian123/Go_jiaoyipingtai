package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/market/mclient"
	"grpc-common/ucenter/ucclient"
	"market-api/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	ExchangeRateRpc mclient.ExchangeRate
	UCLoginRpc      ucclient.Login
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ExchangeRateRpc: mclient.NewExchangeRate(zrpc.MustNewClient(c.MarketRpc)),
	}
}
