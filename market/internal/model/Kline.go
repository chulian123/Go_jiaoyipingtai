package model

import (
	"grpc-common/market/types/market"
	"mscoin-common/op"
)

type Kline struct {
	Period       string  `bson:"period,omitempty" `
	OpenPrice    float64 `bson:"openPrice,omitempty"`
	HighestPrice float64 `bson:"highestPrice,omitempty"`
	LowestPrice  float64 `bson:"lowestPrice,omitempty"`
	ClosePrice   float64 `bson:"closePrice,omitempty"`
	Time         int64   `bson:"time,omitempty"`
	Count        float64 `bson:"count,omitempty"`    //成交笔数
	Volume       float64 `bson:"volume,omitempty"`   //成交量
	Turnover     float64 `bson:"turnover,omitempty"` //成交额
}

func (*Kline) Table(symbol, period string) string {
	return "exchange_kline_" + symbol + "_" + period
}

func (k *Kline) ToCoinThumb(symbol string, end *Kline) *market.CoinThumb {
	// 创建一个CoinThumb对象
	ct := &market.CoinThumb{}
	// 设置CoinThumb对象的Symbol字段为给定的symbol参数
	ct.Symbol = symbol
	// 设置CoinThumb对象的Close字段为K线数据的ClosePrice字段
	ct.Close = k.ClosePrice
	// 设置CoinThumb对象的Open字段为K线数据的OpenPrice字段
	ct.Open = k.OpenPrice
	// 设置CoinThumb对象的Zone字段为0
	ct.Zone = 0
	// 计算CoinThumb对象的Change字段，为当前K线数据的ClosePrice字段与end参数的ClosePrice字段之差
	ct.Change = k.ClosePrice - end.ClosePrice
	// 计算CoinThumb对象的Chg字段，为Change字段除以end参数的LowestPrice字段，并保留5位小数
	ct.Chg = op.MulN(op.DivN(ct.Change, end.ClosePrice, 5), 100, 5)
	// 设置CoinThumb对象的UsdRate字段为K线数据的ClosePrice字段
	ct.UsdRate = k.ClosePrice
	// 设置CoinThumb对象的BaseUsdRate字段为1
	ct.BaseUsdRate = 1
	// 返回转换后的CoinThumb对象

	ct.DateTime = k.Time

	return ct
}

func DefaultCoinThu(symbol string) *market.CoinThumb {
	ct := &market.CoinThumb{}
	ct.Symbol = symbol
	ct.Trend = []float64{}
	return ct
}
