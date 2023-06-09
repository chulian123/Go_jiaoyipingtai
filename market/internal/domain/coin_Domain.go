package domain

import (
	"context"
	"errors"
	"market/internal/dao"
	"market/internal/model"
	"market/internal/repo"
	"mscoin-common/msdb"
)

type CoinDomain struct {
	coinRepo repo.CoinRepo
}

func (d *CoinDomain) FindCoinInfo(ctx context.Context, unit string) (*model.Coin, error) {
	coin, err := d.coinRepo.FindByUnit(ctx, unit)
	if err != nil {
		return nil, err
	}
	if coin == nil {
		return nil, errors.New("不支持的这个coin:" + unit)
	}
	return coin, nil

}

func NewCoinDomain(db *msdb.MsDB) *CoinDomain {
	return &CoinDomain{
		coinRepo: dao.NewCoinDao(db),
	}
}
