package domain

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"grpc-common/market/mclient"
	"mscoin-common/msdb"
	"mscoin-common/msdb/tran"
	"mscoin-common/op"
	"mscoin-common/tools"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberWalletDomain struct {
	memberWalletRepo repo.MemberWalletRepo
	transaction      tran.Transaction
	marketRpc        mclient.Market
	redisCache       cache.Cache
}

func (d *MemberWalletDomain) FindWalletBySymbol(ctx context.Context, id int64, name string, coin *mclient.Coin) (*model.MemberWalletCoin, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, id, name)
	if err != nil {
		return nil, err
	}
	if mw == nil {
		//新建并存储
		mw, walletCoin := model.NewMemberWallet(id, coin)
		err := d.memberWalletRepo.Save(ctx, mw)
		if err != nil {
			return nil, err
		}
		return walletCoin, nil
	}
	nwc := &model.MemberWalletCoin{}
	copier.Copy(nwc, mw)
	nwc.Coin = coin
	return nwc, nil
}

func (d *MemberWalletDomain) Freeze(ctx context.Context, conn msdb.DbConn, userId int64, money float64, symbol string) error {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, userId, symbol)
	if err != nil {
		return err
	}
	if mw.Balance < money {
		return errors.New("余额不足")
	}
	err = d.memberWalletRepo.UpdateFreeze(ctx, conn, userId, symbol, money)
	if err != nil {
		return err
	}
	return nil
}

func (d *MemberWalletDomain) FindWalletByMemIdAndCoin(ctx context.Context, memberId int64, coinName string) (*model.MemberWallet, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinName(ctx, memberId, coinName)
	if err != nil {
		return nil, err
	}
	return mw, nil
}

func (d *MemberWalletDomain) FindWalletByMemIdAndCoinId(ctx context.Context, memberId int64, coinId int64) (*model.MemberWallet, error) {
	mw, err := d.memberWalletRepo.FindByIdAndCoinId(ctx, memberId, coinId)
	if err != nil {
		return nil, err
	}
	return mw, nil
}

//func (d *MemberWalletDomain) UpdateWallet(ctx context.Context, wallet *model.MemberWallet) error {
//	return d.transaction.Action(func(conn msdb.DbConn) error {
//		err := d.memberWalletRepo.UpdateWallet(ctx, conn, wallet.Id, wallet.Balance, wallet.FrozenBalance)
//		if err != nil {
//			return err
//		}
//		return nil
//	})
//}

func (d *MemberWalletDomain) UpdateWalletCoinAndBase(ctx context.Context, baseWallet *model.MemberWallet, coinWallet *model.MemberWallet) error {
	return d.transaction.Action(func(conn msdb.DbConn) error {
		err := d.memberWalletRepo.UpdateWallet(ctx, conn, baseWallet.Id, baseWallet.Balance, baseWallet.FrozenBalance)
		if err != nil {
			return err
		}
		err = d.memberWalletRepo.UpdateWallet(ctx, conn, coinWallet.Id, coinWallet.Balance, coinWallet.FrozenBalance)
		if err != nil {
			return err
		}
		return nil
	})
}

func (d *MemberWalletDomain) FindWallet(ctx context.Context, userId int64) (list []*model.MemberWalletCoin, err error) {
	memberWallets, err := d.memberWalletRepo.FindByMemberId(ctx, userId)
	if err != nil {
		return nil, err
	}

	//查询cny的汇率
	var cnyRateStr string
	d.redisCache.Get("USDT::CNY::RATE", &cnyRateStr)
	var cnyRate float64 = 7
	if cnyRateStr != "" {
		cnyRate = tools.ToFloat64(cnyRateStr)
	}
	//需要查询 币种的详情
	for _, v := range memberWallets {
		coinInfo, err := d.marketRpc.FindCoinInfo(ctx, &mclient.MarketReq{
			Unit: v.CoinName,
		})
		if err != nil {
			return nil, err
		}
		if coinInfo.Unit == "USDT" {
			coinInfo.CnyRate = cnyRate
			coinInfo.UsdRate = 1
		} else {
			var usdtRateStr string
			var usdtRate float64 = 20000
			d.redisCache.Get(v.CoinName+"::USDT::RATE", &usdtRateStr)
			if usdtRateStr != "" {
				usdtRate = tools.ToFloat64(usdtRateStr)
			}
			coinInfo.UsdRate = usdtRate
			coinInfo.CnyRate = op.MulFloor(cnyRate, coinInfo.UsdRate, 10)
		}
		list = append(list, v.Copy(coinInfo))
	}
	return list, nil
}

func (d *MemberWalletDomain) UpdateAddress(ctx context.Context, wallet *model.MemberWallet) error {
	return d.memberWalletRepo.UpdateAddress(ctx, wallet)
}

func (d *MemberWalletDomain) GetAllAddress(ctx context.Context, coinName string) ([]string, error) {
	return d.memberWalletRepo.FindAllAddress(ctx, coinName)
}

func (d *MemberWalletDomain) FindByAddress(ctx context.Context, address string) (*model.MemberWallet, error) {
	return d.memberWalletRepo.FindByAddress(ctx, address)
}

func NewMemberWalletDomain(db *msdb.MsDB, marketRpc mclient.Market, redisCache cache.Cache) *MemberWalletDomain {
	return &MemberWalletDomain{
		memberWalletRepo: dao.NewMemberWalletDao(db),
		transaction:      tran.NewTransaction(db.Conn),
		marketRpc:        marketRpc,
		redisCache:       redisCache,
	}
}
