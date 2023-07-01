package svc

import (
	"jobcenter/internal/config"
	"jobcenter/internal/database"
)

type ServiceContext struct {
	Config      config.Config
	MongoClient *database.MongoClient
	KafkaClient *database.KafkaClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	//启动了Kafka客户端的写入功能。这意味着客户端开始监听要发送到Kafka的数据，并将其写入相应的Kafka主题
	client := database.NewKafkaClient(c.Kafka)
	client.StartWrite()
	return &ServiceContext{
		Config:      c,
		MongoClient: database.ConnectMongo(c.Mongo),
		KafkaClient: client,
	}
}
