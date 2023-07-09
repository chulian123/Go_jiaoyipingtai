package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/svc"
	"market-api/internal/types"
	"time"
)

//市场业务的逻辑代码

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (resp []*types.CoinThumbResp, err error) {
	//symbolThumbTrend, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(context.Background(),
	//	&market.MarketReq{
	//		Ip: req.Ip,
	//	})
	//if err != nil {
	//	return nil, err
	//}
	//if err := copier.Copy(&list, symbolThumbTrend.List); err != nil {
	//	return nil, err
	//}
	//return
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var thumbs []*market.CoinThumb
	thumb := l.svcCtx.Processor.GetThumb()
	isCache := false
	if thumb != nil {
		switch thumb.(type) {
		case []*market.CoinThumb:
			thumbs = thumb.([]*market.CoinThumb)
			isCache = true
		}
	}
	if !isCache {
		thumbResp, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(ctx, &market.MarketReq{
			Ip: req.Ip,
		})
		if err != nil {
			return nil, err
		}
		thumbs = thumbResp.List

	}
	if err := copier.Copy(&resp, thumbs); err != nil {
		return nil, errors.New("数据格式有误")
	}
	for _, v := range resp {
		if v.Trend == nil {
			v.Trend = []float64{}
		}
	}
	return
}

func (l *MarketLogic) SymbolThumb(req *types.MarketReq) (resp []*types.CoinThumbResp, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	var thumbs []*market.CoinThumb
	thumb := l.svcCtx.Processor.GetThumb()
	if thumb != nil {
		switch thumb.(type) {
		case []*market.CoinThumb:
			thumbs = thumb.([]*market.CoinThumb)
		}
	}
	if err := copier.Copy(&resp, thumbs); err != nil {
		return nil, errors.New("数据格式有误")
	}
	for _, v := range resp {
		if v.Trend == nil {
			v.Trend = []float64{}
		}
	}
	return
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
