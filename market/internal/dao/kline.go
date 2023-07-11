package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"market/internal/model"

	"log"
)

type KlineDao struct {
	db *mongo.Database
}

func (k *KlineDao) SaveBatch(ctx context.Context, data []*model.Kline, symbol, period string) error {
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	ds := make([]interface{}, len(data))
	for i, v := range data {
		ds[i] = v
	}
	_, err := collection.InsertMany(ctx, ds)
	return err
}

func (k *KlineDao) DeleteGtTime(ctx context.Context, time int64, symbol string, period string) error {
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	deleteResult, err := collection.DeleteMany(ctx, bson.D{{"time", bson.D{{"$gte", time}}}})
	if err != nil {
		return err
	}
	log.Printf("%s %s 删除了%d条数据 \n", symbol, period, deleteResult.DeletedCount)
	return nil
}

func (k *KlineDao) FindBySymbol(ctx context.Context, symbol, period string, count int64) (list []*model.Kline, err error) {
	//按照时间 降序排序
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	//ctx：上下文对象，用于控制查询的生命周期和取消操作。
	//bson.D{{}}：一个空的BSON文档作为查询条件，表示匹配所有文档。
	//&options.FindOptions{}：用于设置查询选项的结构体。
	//Limit: &count：限制查询结果的数量，count变量指定了查询结果的文档数。
	//Sort: bson.D{{"time", -1}}：指定按照"time"字段进行降序排序。
	cur, err := collection.Find(ctx, bson.D{{}}, &options.FindOptions{
		Limit: &count,
		Sort:  bson.D{{"time", -1}}, //排序操作 降序
	})
	//查找所有文档并按照"time"字段进行降序排序，并限制查询结果的数量。查询结果会存储在cur变量中，err变量用于捕获可能的错误信息。
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}
func (k *KlineDao) FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, s string) (list []*model.Kline, err error) {
	//按照时间范围来查询
	//安装时间范围 查询
	mk := &model.Kline{}
	sortInt := -1
	if "asc" == s {
		sortInt = 1
	}
	collection := k.db.Collection(mk.Table(symbol, period))
	cur, err := collection.Find(ctx,
		bson.D{{"time", bson.D{{"$gte", from}, {"$lte", end}}}},
		&options.FindOptions{
			Sort: bson.D{{"time", sortInt}},
		})
	if err != nil {
		return nil, err
	}

	err = cur.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return
}

func NewKlineDao(db *mongo.Database) *KlineDao {
	return &KlineDao{
		db: db,
	}
}
