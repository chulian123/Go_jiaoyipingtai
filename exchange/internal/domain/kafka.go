package domain

import (
	"context"
	"encoding/json"
	"exchange/internal/database"
	"exchange/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaDomain struct {
	cli         *database.KafkaClient
	orderDomain *ExchangeOrderDomain
}

func (d *KafkaDomain) SendOrderAdd(
	topic string,
	userId int64,
	orderId string,
	money float64,
	symbol string,
	direction int,
	baseSymbol string,
	coinSymbol string) error {
	m := make(map[string]any)
	m["userId"] = userId
	m["orderId"] = orderId
	m["money"] = money
	m["symbol"] = symbol
	m["direction"] = direction
	m["baseSymbol"] = baseSymbol
	m["coinSymbol"] = coinSymbol
	marshal, _ := json.Marshal(m)
	data := database.KafkaData{
		Topic: topic,
		Key:   []byte(orderId),
		Data:  marshal,
	}
	err := d.cli.SendSync(data)
	logx.Info("创建订单，发消息成功,orderId=" + orderId)
	return err
}

type OrderResult struct {
	UserId  int64  `json:"userId"`
	OrderId string `json:"orderId"`
}

func (d *KafkaDomain) WaitAddOrderResult() {
	d.cli.StartRead("exchange_order_init_complete_trading")
	for {
		kafkaData := d.cli.Read()
		logx.Info("读取exchange_order_init_complete_trading 消息成功:" + string(kafkaData.Key))
		var orderResult OrderResult
		json.Unmarshal(kafkaData.Data, &orderResult)
		exchangeOrder, err := d.orderDomain.FindOrderByOrderId(context.Background(), orderResult.OrderId)
		if err != nil {
			logx.Error(err)
			err := d.orderDomain.UpdateStatusCancel(context.Background(), orderResult.OrderId)
			if err != nil {
				d.cli.RPut(kafkaData)
			}
			continue
		}
		if exchangeOrder == nil {
			logx.Error("订单id不存在")
			continue
		}
		if exchangeOrder.Status != model.Init {
			logx.Error("订单已经被处理过")
			continue
		}
		err = d.orderDomain.UpdateOrderStatusTrading(context.Background(), orderResult.OrderId)
		if err != nil {
			logx.Error(err)
			d.cli.RPut(kafkaData)
			continue
		}
	}
}

func NewKafkaDomain(cli *database.KafkaClient, orderDomain *ExchangeOrderDomain) *KafkaDomain {
	k := &KafkaDomain{
		cli:         cli,
		orderDomain: orderDomain,
	}
	go k.WaitAddOrderResult()
	return k
}
