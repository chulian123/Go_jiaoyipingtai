package domain

import (
	"context"
	"grpc-common/market/types/market"
	"market/internal/dao"
	"market/internal/database"
	"market/internal/model"
	"market/internal/repo"
	"mscoin-common/op"
	"mscoin-common/tools"
	"time"
)

type MarketDomain struct {
	klineRepo repo.KlineRepo //接口不要*
}

func (d MarketDomain) SymbolThumbTrend(coins []*model.ExchangeCoin) []*market.CoinThumb {
	//业务模型  ==  rpc传输模型
	coinThumbs := make([]*market.CoinThumb, len(coins))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for i, v := range coins {
		form := tools.ZeroTime()
		end := time.Now().UnixMilli()
		klines, err := d.klineRepo.FindBySymbolTime(ctx, v.Symbol, "1H", form, end, "")
		if err != nil {
			//如果当我们读取不到数据时就 通过工具类给她一个默认值先
			coinThumbs[i] = model.DefaultCoinThu(v.Symbol)
			continue
		}
		length := len(klines)
		if length <= 0 {
			coinThumbs[i] = model.DefaultCoinThu(v.Symbol)
			continue
		}
		//降序排列0 最新数据 length-1 今天最新的数据
		//构建趋势\
		trend := make([]float64, length)
		var low float64 = klines[0].LowestPrice
		var volume float64 = 0
		var high float64 = 0
		var turnover float64 = 0
		for i := length - 1; i >= 0; i-- {
			trend[i] = klines[i].ClosePrice
			highPrice := klines[i].HighestPrice
			if highPrice > high {
				high = highPrice
			}
			lowPrice := klines[i].HighestPrice
			if lowPrice < low {
				high = lowPrice
			}
			volume = op.AddN(volume, klines[i].Volume, 8)
			turnover = op.AddN(volume, klines[i].Turnover, 8)
		}
		newkline := klines[0]
		oldkline := klines[length-1]
		thumb := newkline.ToCoinThumb(v.Symbol, oldkline)
		thumb.Trend = trend
		thumb.High = high
		thumb.Low = low
		thumb.Volume = volume
		thumb.Turnover = turnover
		coinThumbs[i] = thumb
	}
	return coinThumbs
}

func (d *MarketDomain) HistoryKline(ctx context.Context, symbol string, from int64, to int64, period string) ([]*market.History, error) {
	klines, err := d.klineRepo.FindBySymbolTime(ctx, symbol, period, from, to, "asc")
	if err != nil {
		return nil, err
	}
	list := make([]*market.History, len(klines))
	for i, v := range klines {
		h := &market.History{}
		h.Time = v.Time
		h.Open = v.OpenPrice
		h.High = v.HighestPrice
		h.Low = v.LowestPrice
		h.Volume = v.Volume
		h.Close = v.ClosePrice
		list[i] = h
	}
	return list, nil
}

func NewMarketDomain(mongoClient *database.MongoClient) *MarketDomain {
	return &MarketDomain{
		klineRepo: dao.NewKlineDao(mongoClient.Db),
	}
}
