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
	exchangeCoinRepo repo.ExchangeCoinRepo
}

func NewExchangeCoinDomain(db *msdb.MsDB) *ExchangeCoinDomain {
	return &ExchangeCoinDomain{
		exchangeCoinRepo: dao.NewExchangeCoinDao(db),
	}
}

func (d *ExchangeCoinDomain) FindVisible(ctx context.Context) []*model.ExchangeCoin {
	list, err := d.exchangeCoinRepo.FindVisible(ctx)
	if err != nil {
		logx.Error(err)
		return []*model.ExchangeCoin{}
	}
	return list
}

func (d *ExchangeCoinDomain) FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error) {
	exchangeCoin, err := d.exchangeCoinRepo.FindBySymbol(ctx, symbol)
	if err != nil {
		return nil, err
	}
	if exchangeCoin == nil {
		return nil, errors.New("交易对不存在")
	}
	return exchangeCoin, nil
}
