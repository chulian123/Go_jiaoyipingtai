package domain

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/repo"
	"mscoin-common/msdb"
)

type ExchangeCoinDomain struct {
	ExchangeCoinRepo repo.ExchangeCoinRepo
}

func NewExchangeCoinDomain(db *msdb.MsDB) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		ExchangeCoinRepo: dao.NewExchangeCoinDao(db),
	}
}

// 查询方式
func (d *ExchangeCoinDomain) FindVisible(ctx context.Context) []*model.ExchangeCoin {
	list, err := d.ExchangeCoinRepo.FindVisible(ctx)
	if err != nil {
		logx.Error(err)
		return []*model.ExchangeCoin{}
	}
	return list
}

func (d *ExchangeCoinDomain) FindByFindSymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	exChangeCoin, err := d.ExchangeCoinRepo.FindByFindSymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}
	if exChangeCoin == nil {
		return nil, errors.New("交易对不存在")
	}
	return exChangeCoin, err
}
