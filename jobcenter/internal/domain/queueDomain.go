package domain

import (
	"encoding/json"
	"jobcenter/internal/database"
	"jobcenter/internal/model"
	"log"
)

// fafka Domain
const KLINE1M = "kline_1m"

type QueueDomain struct {
	kafkaCli *database.KafkaClient
}

func (d *QueueDomain) Send1mKline(data []string, symbol string) {
	kline := model.NewKline(data, "1m") // 创建一个 Kline 对象，表示 K 线数据，使用 1 分钟的时间间隔
	bytes, _ := json.Marshal(kline)     // 将 Kline 对象序列化为 JSON 字节数组
	msg := database.KafkaData{
		Topic: KLINE1M,        // Kafka 主题，表示数据应该发送到哪个主题中
		Data:  bytes,          // 序列化后的 Kline 数据
		Key:   []byte(symbol), // 消息的键，使用 symbol 字符串的字节数组
	}
	d.kafkaCli.Send(msg) // 将消息发送到 Kafka 队列中进行处理

	log.Println("=================发送数据成功==============")
}

func NewQueueDomain(kafkaCli *database.KafkaClient) *QueueDomain {
	return &QueueDomain{
		kafkaCli: kafkaCli,
	}
}
