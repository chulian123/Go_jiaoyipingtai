package processor

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"grpc-common/market/types/market"
	"market-api/internal/database"
	"market-api/internal/model"
	"sync"
	"time"
)

// 专门处理k线数据的文件
const KLINE1M = "kline_1m"
const KLINE = "kline"
const TRADE = "trade"

type ProcessData struct {
	Type string //trade 交易 kline k线
	Key  []byte
	Data []byte
}

type MarketHandler interface {
	HandleTrade(symbol string, data []byte)                                               //处理交易数据
	HandleKLine(symbol string, kline *model.Kline, thumpMap map[string]*market.CoinThumb) //处理k线数据
}

// 处理器接口
type Processor interface {
	GetThumb() any
	Process(data ProcessData)
	AddHandler(h MarketHandler)
}

// 实现类
type DefaultProcessor struct {
	KafkaCli *database.KafkaClient
	handlers []MarketHandler
	thumpMap map[string]*market.CoinThumb
	rwLock   sync.RWMutex //一个读写锁成员变量
}

func NewDefaultProcessor(KafkaCli *database.KafkaClient) *DefaultProcessor {
	return &DefaultProcessor{
		KafkaCli: KafkaCli,
		handlers: make([]MarketHandler, 0),
		thumpMap: make(map[string]*market.CoinThumb),
	}
}

func (d *DefaultProcessor) Process(data ProcessData) {

	if data.Type == KLINE {
		symbol := string(data.Key)
		kline := &model.Kline{}
		json.Unmarshal(data.Data, kline)
		for _, v := range d.handlers {
			v.HandleKLine(symbol, kline, d.thumpMap)
		}
	}
}

func (d *DefaultProcessor) AddHandler(h MarketHandler) {
	//发送到websocket服务里
	d.handlers = append(d.handlers, h)

}

func (d *DefaultProcessor) GetThumb() any {
	//GetThumb() 方法中使用读写锁来保护 thumpMap 的读写操作
	d.rwLock.RLock()         // 获取读锁
	defer d.rwLock.RUnlock() // 在函数返回时释放读锁
	cs := make([]*market.CoinThumb, len(d.thumpMap))
	i := 0
	for _, v := range d.thumpMap {
		cs[i] = v
		i++
	}
	return cs
}

// 初始化的操作
func (p *DefaultProcessor) Init(marketRpc mclient.Market) {
	//在kafka里面拿数据
	//p.AddHandler()
	//接收kline 1m的同步数据
	p.startReadFromKafka(KLINE1M, KLINE)
	p.initThumbMap(marketRpc)
}

func (p *DefaultProcessor) startReadFromKafka(topic string, tp string) {
	//开始从topic里面读取数据
	//要先start 然后在read
	p.KafkaCli.StartRead(topic)
	go p.dealQueueData(p.KafkaCli, tp)

}

func (p *DefaultProcessor) dealQueueData(cli *database.KafkaClient, tp string) {
	//队列的数据
	msg := cli.Read()
	data := ProcessData{
		Type: tp,
		Key:  msg.Key,
		Data: msg.Data,
	}
	//k线数据
	p.Process(data)
}

func (d *DefaultProcessor) initThumbMap(marketRpc mclient.Market) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	thumbResp, err := marketRpc.FindSymbolThumbTrend(ctx, &market.MarketReq{})
	if err != nil {
		logx.Info(err)
	} else {
		list := thumbResp.List
		for _, v := range list {
			d.thumpMap[v.Symbol] = v
		}
	}
}
