package handler

import (
	"market-api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	//封装一下的url
	//如果需要加中间件 怎么办
	//注册url
	rate := NewExchangeRateHandler(serverCtx)
	rateGroup := r.Group()
	rateGroup.Post("/exchange-rate/usd/:unit", rate.UsdRate)

	market := NewMarketHandler(serverCtx)
	marketGroup := r.Group()
	marketGroup.Post("/symbol-thumb-trend", market.SymbolThumbTrend)

	wsGroup := r.Group()
	wsGroup.GetNoPrefix("/socket.io", nil)
	wsGroup.PostNoPrefix("/socket.io", nil)
}
