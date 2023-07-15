package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"market/internal/model"
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
	//按照时间 降序排列
	mk := &model.Kline{}
	collection := k.db.Collection(mk.Table(symbol, period))
	cur, err := collection.Find(ctx, bson.D{{}}, &options.FindOptions{
		Limit: &count,
		Sort:  bson.D{{"time", -1}},
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
func (k *KlineDao) FindBySymbolTime(ctx context.Context, symbol, period string, from, end int64, sort string) (list []*model.Kline, err error) {
	//安装时间范围 查询
	mk := &model.Kline{}
	sortInt := -1
	if "asc" == sort {
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
