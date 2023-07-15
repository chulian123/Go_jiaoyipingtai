package model

import (
	"github.com/jinzhu/copier"
	"mscoin-common/enum"
)

type ExchangeOrder struct {
	Id            int64   `gorm:"column:id"`
	OrderId       string  `gorm:"column:order_id"`
	Amount        float64 `gorm:"column:amount"`
	BaseSymbol    string  `gorm:"column:base_symbol"`
	CanceledTime  int64   `gorm:"column:canceled_time"`
	CoinSymbol    string  `gorm:"column:coin_symbol"`
	CompletedTime int64   `gorm:"column:completed_time"`
	Direction     int     `gorm:"column:direction"`
	MemberId      int64   `gorm:"column:member_id"`
	Price         float64 `gorm:"column:price"`
	Status        int     `gorm:"column:status"`
	Symbol        string  `gorm:"column:symbol"`
	Time          int64   `gorm:"column:time"`
	TradedAmount  float64 `gorm:"column:traded_amount"`
	Turnover      float64 `gorm:"column:turnover"`
	Type          int     `gorm:"column:type"`
	UseDiscount   string  `gorm:"column:use_discount"`
}

func (*ExchangeOrder) TableName() string {
	return "exchange_order"
}

// status
const (
	Trading = iota
	Completed
	Canceled
	OverTimed
	Init
)

var StatusMap = enum.Enum{
	Trading:   "TRADING",
	Completed: "COMPLETED",
	Canceled:  "CANCELED",
	OverTimed: "OVERTIMED",
}

// direction
const (
	BUY = iota
	SELL
)

var DirectionMap = enum.Enum{
	BUY:  "BUY",
	SELL: "SELL",
}

// type
const (
	MarketPrice = iota
	LimitPrice
)

var TypeMap = enum.Enum{
	MarketPrice: "MARKET_PRICE",
	LimitPrice:  "LIMIT_PRICE",
}

type ExchangeOrderVo struct {
	OrderId       string  `gorm:"column:order_id"`
	Amount        float64 `gorm:"column:amount"`
	BaseSymbol    string  `gorm:"column:base_symbol"`
	CanceledTime  int64   `gorm:"column:canceled_time"`
	CoinSymbol    string  `gorm:"column:coin_symbol"`
	CompletedTime int64   `gorm:"column:completed_time"`
	Direction     string  `gorm:"column:direction"`
	MemberId      int64   `gorm:"column:member_id"`
	Price         float64 `gorm:"column:price"`
	Status        string  `gorm:"column:status"`
	Symbol        string  `gorm:"column:symbol"`
	Time          int64   `gorm:"column:time"`
	TradedAmount  float64 `gorm:"column:traded_amount"`
	Turnover      float64 `gorm:"column:turnover"`
	Type          string  `gorm:"column:type"`
	UseDiscount   string  `gorm:"column:use_discount"`
}

func (old *ExchangeOrder) ToVo() *ExchangeOrderVo {
	eo := &ExchangeOrderVo{}
	copier.Copy(eo, old)
	eo.Status = StatusMap.Value(old.Status)
	eo.Direction = DirectionMap.Value(old.Direction)
	eo.Type = TypeMap.Value(old.Type)
	return eo
}

func NewOrder() *ExchangeOrder {
	return &ExchangeOrder{}
}
