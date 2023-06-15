package logic

import (
	"context"
	"grpc-common/market/types/market"

	"market/internal/domain"
	"market/internal/svc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type MarketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeCoinDomain *domain.ExchangeCoinDomain
	marketDomain       *domain.MarketDomain
}

func (l *MarketLogic) FindSymbolThumbTrend(in *market.MarketReq) (*market.SymbolThumbRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exchangeCoins := l.exchangeCoinDomain.FindVisible(ctx)
	//在查询mongo相对应的数据
	//查询一个小时间隔的 可以根据时间间隔来进行查询  当天价格变化趋势
	coinThumbs := l.marketDomain.SymbolThumbTrend(exchangeCoins)

	//假数据的部分操作
	//coinThumbs := make([]*market.CoinThumb, len(exchangeCoins))
	//for i, v := range exchangeCoins {
	//	trend := make([]float64, 0)
	//	for p := 0; p <= 100; p++ {
	//		trend = append(trend, rand.Float64())
	//	}
	//	ct := &market.CoinThumb{}
	//	ct.Symbol = v.Symbol
	//	ct.Trend = trend
	//	coinThumbs[i] = ct
	//}
	return &market.SymbolThumbRes{
		List: coinThumbs,
	}, nil
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeCoinDomain: domain.NewExchangeCoinDomain(svcCtx.Db),
		marketDomain:       domain.NewMarketDomain(svcCtx.MongoClient),
	}
}
