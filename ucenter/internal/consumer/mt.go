package consumer

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"mscoin-common/msdb"
	"time"
	"ucenter/internal/database"
	"ucenter/internal/domain"
)

type BitCoinTransactionResult struct {
	Value  float64 `json:"value"`
	Time   int64   `json:"time"`
	Adress string  `json:"adress"`
	Type   string  `json:"type"`
	Symbol string  `json:"symbol"`
}

func BitCoinTransaction(redisClient *redis.Redis, kakfaClient *database.KafkaClient, db *msdb.MsDB) {
	for {
		kafkaData := kakfaClient.Read()
		var result BitCoinTransactionResult
		json.Unmarshal(kafkaData.Data, &result)
		//解析出来数据 调用domian然后存储到数据库
		transactionDomain := domain.NewMemberTransactionDomain(db)
		err := transactionDomain.SaveRecharge(result.Adress, result.Value, result.Time, result.Type, result.Symbol)
		if err != nil {
			kakfaClient.Rput(kafkaData)
			time.Sleep(200 * time.Millisecond)
		}

	}
}
