package repo

import (
	"context"
	"market/internal/model"
)

//接口 对应Member的操作

type ExchangeCoinRepo interface {
	FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error)
	FindByFindSymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) //查询信息接口
}
