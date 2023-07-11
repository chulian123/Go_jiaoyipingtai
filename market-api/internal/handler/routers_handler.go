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
	marketGroup.Post("/symbol-thumb", market.SymbolThumb)
	marketGroup.Post("/symbol-info", market.SymbolInfo)
	marketGroup.Post("/coin-info", market.CoinInfo)
	marketGroup.Get("/history", market.History)

	wsGroup := r.Group()
	wsGroup.GetNoPrefix("/socket.io", nil)
	wsGroup.PostNoPrefix("/socket.io", nil)
}
