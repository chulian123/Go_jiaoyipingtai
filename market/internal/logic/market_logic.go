package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market/internal/domain"
	"market/internal/svc"
	"time"
)

type MarketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeCoinDomain *domain.ExchangeCoinDomain
	marketDomain       *domain.MarketDomain
	coinDomain         *domain.CoinDomain
}

func (l *MarketLogic) FindSymbolThumbTrend(req *market.MarketReq) (*market.SymbolThumbRes, error) {
	coins := l.exchangeCoinDomain.FindVisible(l.ctx)
	//查询mongo中相应的数据
	//查询1H间隔的 可以根据时间来进行查询 当天的价格变化趋势
	coinThumbs := l.marketDomain.SymbolThumbTrend(coins)
	//coinThumbs := make([]*market.CoinThumb, len(coins))
	//for i, v := range coins {
	//	ct := &market.CoinThumb{}
	//	ct.Symbol = v.Symbol
	//	trend := make([]float64, 0)
	//	for p := 0; p <= 24; p++ {
	//		trend = append(trend, rand.Float64())
	//	}
	//	ct.Trend = trend
	//	coinThumbs[i] = ct
	//}
	return &market.SymbolThumbRes{
		List: coinThumbs,
	}, nil
}

func (l *MarketLogic) FindSymbolInfo(req *market.MarketReq) (*market.ExchangeCoin, error) {
	exchangeCoin, err := l.exchangeCoinDomain.FindBySymbol(l.ctx, req.Symbol)
	if err != nil {
		return nil, err
	}
	ec := &market.ExchangeCoin{}
	copier.Copy(ec, exchangeCoin)
	return ec, nil
}

func (l *MarketLogic) FindCoinInfo(req *market.MarketReq) (*market.Coin, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	coin, err := l.coinDomain.FindCoinInfo(ctx, req.Unit)
	if err != nil {
		return nil, err
	}
	mc := &market.Coin{}
	if err := copier.Copy(mc, coin); err != nil {
		return nil, err
	}
	return mc, nil
}

func (l *MarketLogic) HistoryKline(req *market.MarketReq) (*market.HistoryRes, error) {
	//去mongo 表查询数据 按照时间范围进行查询 同时要排序 按照时间升序
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	period := "1H"
	if req.Resolution == "60" {
		period = "1H"
	} else if req.Resolution == "30" {
		period = "30m"
	} else if req.Resolution == "15" {
		period = "15m"
	} else if req.Resolution == "5" {
		period = "5m"
	} else if req.Resolution == "1" {
		period = "1m"
	} else if req.Resolution == "1D" {
		period = "1D"
	} else if req.Resolution == "1W" {
		period = "1W"
	} else if req.Resolution == "1M" {
		period = "1M"
	}
	histories, err := l.marketDomain.HistoryKline(ctx, req.Symbol, req.From, req.To, period)
	if err != nil {
		return nil, err
	}
	return &market.HistoryRes{
		List: histories,
	}, nil
}

func (l *MarketLogic) FindExchangeCoinVisible(req *market.MarketReq) (*market.ExchangeCoinRes, error) {
	exchangeCoins := l.exchangeCoinDomain.FindVisible(l.ctx)
	var list []*market.ExchangeCoin
	copier.Copy(&list, exchangeCoins)
	return &market.ExchangeCoinRes{
		List: list,
	}, nil
}

func (l *MarketLogic) FindAllCoin(req *market.MarketReq) (*market.CoinList, error) {
	coinlist, err := l.coinDomain.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	var list []*market.Coin
	copier.Copy(&list, coinlist)
	return &market.CoinList{
		List: list,
	}, nil

}

func (l *MarketLogic) FindById(req *market.MarketReq) (*market.Coin, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 5*time.Second)
	defer cancel()
	coin, err := l.coinDomain.FindCoinId(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	mc := &market.Coin{}
	if err := copier.Copy(mc, coin); err != nil {
		return nil, err
	}
	return mc, nil
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeCoinDomain: domain.NewExchangeCoinDomain(svcCtx.Db),
		marketDomain:       domain.NewMarketDomain(svcCtx.MongoClient),
		coinDomain:         domain.NewCoinDomain(svcCtx.Db),
	}
}
