package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"grpc-common/exchange/eclient"
	"grpc-common/exchange/types/order"
	"mscoin-common/msdb"
	"mscoin-common/msdb/tran"
	"time"
	"ucenter/internal/database"
	"ucenter/internal/domain"
)

type OrderAdd struct {
	UserId     int64   `json:"userId"`
	OrderId    string  `json:"orderId"`
	Money      float64 `json:"money"`
	Symbol     string  `json:"symbol"`
	Direction  int     `json:"direction"`
	BaseSymbol string  `json:"baseSymbol"`
	CoinSymbol string  `json:"coinSymbol"`
}

func ExchangeOrderAdd(redisCli *redis.Redis, cli *database.KafkaClient, orderRpc eclient.Order, db *msdb.MsDB) {
	//exchange rpc kafka cli
	for {
		kafkaData := cli.Read()
		orderId := string(kafkaData.Key)
		logx.Info("接收到创建订单的消息,orderId=" + orderId)
		var orderAdd OrderAdd
		json.Unmarshal(kafkaData.Data, &orderAdd)
		if orderId != orderAdd.OrderId {
			logx.Error("消息数据有误")
			continue
		}
		ctx := context.Background()
		//冻结钱
		exchangeOrderOrigin, err := orderRpc.FindByOrderId(ctx, &order.OrderReq{
			OrderId: orderId,
		})
		if err != nil {
			logx.Error(err)
			cancelOrder(ctx, kafkaData, orderId, orderRpc, cli)
			continue
		}
		if exchangeOrderOrigin == nil {
			logx.Error("orderId :" + orderId + " 不存在")
			continue
		}
		//4 init状态
		if exchangeOrderOrigin.Status != 4 {
			logx.Error("orderId :" + orderId + " 已经被操作过了")
			continue
		}
		lock := redis.NewRedisLock(redisCli, "exchange_order::"+fmt.Sprintf("%d::%s", orderAdd.UserId, orderId))
		acquire, err := lock.Acquire()
		if err != nil {
			logx.Error(err)
			logx.Info("已经有别的进程在处理了....")
			continue
		}
		if acquire {
			transaction := tran.NewTransaction(db.Conn)
			walletDomain := domain.NewMemberWalletDomain(db)
			err := transaction.Action(func(conn msdb.DbConn) error {
				if orderAdd.Direction == 0 {
					//buy
					err := walletDomain.Freeze(ctx, conn, orderAdd.UserId, orderAdd.Money, orderAdd.BaseSymbol)
					return err
				} else {
					err := walletDomain.Freeze(ctx, conn, orderAdd.UserId, orderAdd.Money, orderAdd.CoinSymbol)
					return err
				}
			})
			if err != nil {
				cancelOrder(ctx, kafkaData, orderId, orderRpc, cli)
				continue
			}
			//需要将状态 改为trading
			//都完成后 通知订单进行状态变更 需要保证一定发送成功
			for {
				m := make(map[string]any)
				m["userId"] = orderAdd.UserId
				m["orderId"] = orderId
				marshal, _ := json.Marshal(m)
				data := database.KafkaData{
					Topic: "exchange_order_init_complete_trading",
					Key:   []byte(orderId),
					Data:  marshal,
				}
				err := cli.SendSync(data)
				if err != nil {
					logx.Error(err)
					time.Sleep(250 * time.Millisecond)
					continue
				}
				logx.Info("发送exchange_order_init_complete_trading 消息成功:" + orderId)
				break
			}
			lock.Release()
		}
	}
}

func cancelOrder(ctx context.Context, data database.KafkaData, orderId string, orderRpc eclient.Order, cli *database.KafkaClient) {
	_, err := orderRpc.CancelOrder(ctx, &order.OrderReq{
		OrderId: orderId,
	})
	if err != nil {
		cli.Rput(data)
	}
}
