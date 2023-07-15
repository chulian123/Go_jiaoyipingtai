package processor

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
	"market-api/internal/database"
	"market-api/internal/model"
)

const KLINE1M = "kline_1m"
const KLINE = "kline"
const TRADE = "trade"

type MarketHandler interface {
	HandleTrade(symbol string, data []byte)
	HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*market.CoinThumb)
}

type ProcessData struct {
	Type string //trade 交易 kline k线
	Key  []byte
	Data []byte
}

type Processor interface {
	GetThumb() any
	Process(data ProcessData)
	AddHandler(h MarketHandler)
}

type DefaultProcessor struct {
	kafkaCli *database.KafkaClient
	handlers []MarketHandler
	thumbMap map[string]*market.CoinThumb
}

func NewDefaultProcessor(kafkaCli *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		kafkaCli: kafkaCli,
		handlers: make([]MarketHandler, 0),
		thumbMap: make(map[string]*market.CoinThumb),
	}
}

func (d *DefaultProcessor) Process(data ProcessData) {
	if data.Type == KLINE {
		symbol := string(data.Key)
		kline := &model.Kline{}
		json.Unmarshal(data.Data, kline)
		for _, v := range d.handlers {
			v.HandleKLine(symbol, kline, d.thumbMap)
		}
	}
}

func (d *DefaultProcessor) AddHandler(h MarketHandler) {
	//发送到websocket的服务
	d.handlers = append(d.handlers, h)
}

func (p *DefaultProcessor) Init(marketRpc mclient.Market) {
	p.startReadFromKafka(KLINE1M, KLINE)
	p.initThumbMap(marketRpc)
}
func (d *DefaultProcessor) GetThumb() any {
	cs := make([]*market.CoinThumb, len(d.thumbMap))
	i := 0
	for _, v := range d.thumbMap {
		cs[i] = v
		i++
	}
	return cs
}

func (p *DefaultProcessor) startReadFromKafka(topic string, tp string) {
	//一定要先start 后read
	p.kafkaCli.StartRead(topic)
	go p.dealQueueData(p.kafkaCli, tp)
}

func (p *DefaultProcessor) dealQueueData(cli *database.KafkaClient, tp string) {
	//这就是队列的数据
	for {
		msg := cli.Read()
		data := ProcessData{
			Type: tp,
			Key:  msg.Key,
			Data: msg.Data,
		}
		p.Process(data)
	}

}

func (d *DefaultProcessor) initThumbMap(marketRpc mclient.Market) {
	symbolThumbRes, err := marketRpc.FindSymbolThumbTrend(context.Background(),
		&market.MarketReq{})
	if err != nil {
		logx.Info(err)
	} else {
		coinThumbs := symbolThumbRes.List
		for _, v := range coinThumbs {
			d.thumbMap[v.Symbol] = v
		}
	}
}
