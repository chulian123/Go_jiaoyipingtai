package domain

import (
	"encoding/json"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
	"log"
)

const KLINE1M = "kline_1m"
const BtcTransactionTopic = "BTC_TRANSACTION"

type QueueDomain struct {
	kafkaCli *database.KafkaClient
}

func (d *QueueDomain) Send1mKline(data []string, symbol string) {
	kline := model.NewKline(data, "1m")
	bytes, _ := json.Marshal(kline)
	msg := database.KafkaData{
		Topic: KLINE1M,
		Data:  bytes,
		Key:   []byte(symbol),
	}
	d.kafkaCli.Send(msg)
	log.Println("=================Send1mKline发送数据成功==============")
}

func (d *QueueDomain) SendRechar(value float64, address string, time int64) {
	data := make(map[string]any)
	data["value"] = value
	data["address"] = address
	data["time"] = time
	data["type"] = model.RECHARGE
	data["Symbol"] = "BTC"
	marshal, _ := json.Marshal(data)
	msg := database.KafkaData{
		Topic: BtcTransactionTopic,
		Data:  marshal,
		Key:   []byte(address),
	}
	d.kafkaCli.Send(msg)
	log.Println("=================SendRechar发送数据成功==============")
}

func NewQueueDomain(kafkaCli *database.KafkaClient) *QueueDomain {
	return &QueueDomain{
		kafkaCli: kafkaCli,
	}
}
