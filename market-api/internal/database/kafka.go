package database

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"sync"
	"time"
)

// KafkaConfig 定义Kafka的配置参数
type KafkaConfig struct {
	Addr          string `json:"Addr,optional"`
	WriteCap      int    `json:"WriteCap,optional"`
	ReadCap       int    `json:"ReadCap,optional"`
	ConsumerGroup string `json:"ConsumerGroup,optional"`
}

// KafkaData 定义Kafka消息的结构
type KafkaData struct {
	Topic string
	Key   []byte
	Data  []byte
}

// KafkaClient 定义Kafka客户端结构
type KafkaClient struct {
	w         *kafka.Writer
	r         *kafka.Reader
	readChan  chan KafkaData
	writeChan chan KafkaData
	c         KafkaConfig
	closed    bool
	mutex     sync.Mutex
}

// NewKafkaClient 创建一个新的Kafka客户端实例
func NewKafkaClient(c KafkaConfig) *KafkaClient {
	return &KafkaClient{
		c: c,
	}
}

// StartWrite 启动Kafka消息发送
func (k *KafkaClient) StartWrite() {
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.c.Addr),
		Balancer: &kafka.LeastBytes{},
	}
	k.w = w
	k.writeChan = make(chan KafkaData, k.c.WriteCap) //有缓冲的channel
	go k.sendKafka()
}

// Send 向Kafka发送消息
func (w *KafkaClient) Send(data KafkaData) {
	defer func() {
		if err := recover(); err != nil {
			w.closed = true
		}
	}()

	w.writeChan <- data
	w.closed = false
}

// Close 关闭Kafka客户端连接
func (w *KafkaClient) Close() {
	if w.w != nil {
		w.w.Close()
		w.mutex.Lock()
		defer w.mutex.Unlock()
		if !w.closed {
			close(w.writeChan)
			w.closed = true
		}
	}
	if w.r != nil {
		w.r.Close()
	}
}

// sendKafka 发送消息到Kafka
func (w *KafkaClient) sendKafka() {
	for {
		select {
		case data := <-w.writeChan:
			messages := []kafka.Message{
				{
					Topic: data.Topic,
					Key:   data.Key,
					Value: data.Data,
				},
			}
			var err error
			const retries = 3
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			success := false
			for i := 0; i < retries; i++ {
				// 尝试在发布消息之前创建主题
				err = w.w.WriteMessages(ctx, messages...)
				if err == nil {
					success = true
					break
				}
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					success = false
					continue
				}
				if err != nil {
					success = false
					log.Printf("kafka send writemessage err %s \n", err.Error())
				}
			}
			if !success {
				// 如果发送失败，重新放入队列等待消费
				w.Send(data)
			}
		}
	}
}

// StartRead 启动Kafka消息接收
func (k *KafkaClient) StartRead(topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.c.Addr},
		Topic:    topic,
		GroupID:  k.c.ConsumerGroup,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	k.r = r
	k.readChan = make(chan KafkaData, k.c.ReadCap)
	go k.readMsg()
}

// readMsg 从Kafka中读取消息
func (k *KafkaClient) readMsg() {
	for {
		m, err := k.r.ReadMessage(context.Background())
		if err != nil {
			logx.Error(err)
			continue
		}
		data := KafkaData{
			Key:  m.Key,
			Data: m.Value,
		}
		k.readChan <- data
	}
}

func (k *KafkaClient) Read() KafkaData {
	msg := <-k.readChan
	return msg

}

//
//func main() {
//	// 使用示例
//	kafkaConfig := KafkaConfig{
//		Addr:          "localhost:9092",
//		WriteCap:      100,
//		ReadCap:       100,
//		ConsumerGroup: "my-group",
//	}
//	client := NewKafkaClient(kafkaConfig)
//	client.StartWrite()
//	client.StartRead()
//
//	// 发送消息到Kafka
//	data := KafkaData{
//		Topic: "my-topic",
//		Key:   []byte("key"),
//		Data:  []byte("Hello, Kafka!"),
//	}
//	client.Send(data)
//
//	// 接收Kafka消息
//	for {
//		select {
//		case msg := <-client.readChan:
//			log.Printf("Received message: topic=%s, key=%s, data=%s\n", msg.Topic, string(msg.Key), string(msg.Data))
//		}
//	}
//}
