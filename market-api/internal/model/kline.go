package model

import (
	"grpc-common/market/types/market"
	"mscoin-common/op"
)

type Kline struct {
	Period       string  `bson:"period,omitempty"`
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Count        float64 `bson:"count,omitempty"`    //成交笔数
	Volume       float64 `bson:"volume,omitempty"`   //成交量
	Turnover     float64 `bson:"turnover,omitempty"` //成交额
}

func (k *Kline) ToCoinThumb(symbol string, ct *market.CoinThumb) *market.CoinThumb {
	isSame := false
	if ct.Symbol == symbol && ct.DateTime == k.Time {
		//认为是同一个数据
		isSame = true
	}
	if !isSame {
		ct.Close = k.ClosePrice
		ct.Open = k.OpenPrice
		if ct.High < k.HighestPrice {
			ct.High = k.HighestPrice
		}
		ct.Low = k.LowestPrice
		if ct.Low > k.LowestPrice {
			ct.Low = k.LowestPrice
		}
		ct.Zone = 0
		ct.Volume = op.AddN(k.Volume, ct.Volume, 8)
		ct.Turnover = op.AddN(k.Turnover, ct.Turnover, 8)
		ct.Change = k.LowestPrice - ct.Close
		ct.Chg = op.MulN(op.DivN(ct.Change, ct.Close, 5), 100, 3)
		ct.UsdRate = k.ClosePrice
		ct.BaseUsdRate = 1
		ct.Trend = append(ct.Trend, k.ClosePrice)
		ct.DateTime = k.Time
	}
	return ct
}

func (k *Kline) InitCoinThumb(symbol string) *market.CoinThumb {
	ct := &market.CoinThumb{}
	ct.Symbol = symbol
	ct.Close = k.ClosePrice
	ct.Open = k.OpenPrice
	ct.High = k.HighestPrice
	ct.Volume = k.Volume
	ct.Turnover = k.Turnover
	ct.Low = k.LowestPrice
	ct.Zone = 0
	ct.Change = 0
	ct.Chg = 0.0
	ct.UsdRate = k.ClosePrice
	ct.BaseUsdRate = 1
	ct.Trend = make([]float64, 0)

	return ct
}
