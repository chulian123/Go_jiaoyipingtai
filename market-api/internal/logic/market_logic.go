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
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
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

func (l *MarketLogic) SymbolInfo(req types.MarketReq) (resp []*types.ExchangeCoinResp, err error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	esResp, err := l.svcCtx.MarketRpc.FindSymbolInfo(ctx, &market.MarketReq{
		Ip:     req.Ip,
		Symbol: req.Symbol,
	})
	if err != nil {
		return nil, err
	}
	//resp = &types.ExchangeCoinResp{}
	if err := copier.Copy(&resp, esResp); err != nil {
		return nil, err
	}

	return
}

func (l *MarketLogic) CoinInfo(req *types.MarketReq) (*types.Coin, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	coin, err := l.svcCtx.MarketRpc.FindCoinInfo(ctx, &market.MarketReq{
		Unit: req.Unit,
	})
	if err != nil {
		return nil, err
	}
	ec := &types.Coin{}
	if err := copier.Copy(&ec, coin); err != nil {
		return nil, errors.New("数据格式有误")
	}
	return ec, nil
}

func (l *MarketLogic) History(req *types.MarketReq) (*types.HistoryKline, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	historyKline, err := l.svcCtx.MarketRpc.HistoryKline(ctx, &market.MarketReq{
		Symbol:     req.Symbol,
		From:       req.From,
		To:         req.To,
		Resolution: req.Resolution,
	})
	if err != nil {
		return nil, err
	}
	histories := historyKline.List
	var list = make([][]any, len(histories))
	for i, v := range histories {
		content := make([]any, 6)
		content[0] = v.Time
		content[1] = v.Open
		content[2] = v.High
		content[3] = v.Low
		content[4] = v.Close
		content[5] = v.Volume
		list[i] = content
	}
	return &types.HistoryKline{
		List: list,
	}, nil
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
