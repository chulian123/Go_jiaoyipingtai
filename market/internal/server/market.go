// Code genemarketd by goctl. DO NOT EDIT.
// Source: register.proto

package server

import (
	"context"
	"grpc-common/market/types/market"
	"market/internal/logic"
	"market/internal/svc"
)

type MarketServer struct {
	svcCtx *svc.ServiceContext
	market.UnimplementedMarketServer
}

func (e *MarketServer) FindSymbolThumbTrend(ctx context.Context, req *market.MarketReq) (*market.SymbolThumbRes, error) {
	l := logic.NewMarketLogic(ctx, e.svcCtx)
	return l.FindSymbolThumbTrend(req)
}

func (e *MarketServer) FindSymbolInfo(ctx context.Context, req *market.MarketReq) (*market.ExchangeCoin, error) {
	l := logic.NewMarketLogic(ctx, e.svcCtx)
	return l.FindSymbolInfo(req)
}

func (e *MarketServer) FindCoinInfo(ctx context.Context, req *market.MarketReq) (*market.Coin, error) {
	l := logic.NewMarketLogic(ctx, e.svcCtx)
	return l.FindCoinInfo(req)
}
func (e *MarketServer) HistoryKline(ctx context.Context, req *market.MarketReq) (*market.HistoryRes, error) {
	l := logic.NewMarketLogic(ctx, e.svcCtx)
	return l.HistoryKline(req)
}
func NewMarketServer(svcCtx *svc.ServiceContext) *MarketServer {
	return &MarketServer{
		svcCtx: svcCtx,
	}
}
