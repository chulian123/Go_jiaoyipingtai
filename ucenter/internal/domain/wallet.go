package domain

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/mclient"
	"mscoin-common/msdb"
	"mscoin-common/msdb/tran"
	"mscoin-common/op"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"
)

type MemberWalletDomain struct {
	memberWalletRepo repo.MemberWalletRepo
	transaction      tran.Transaction
	marketRpc        mclient.Market
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
	//查询 币种详情
	memberWallet, err := d.memberWalletRepo.FindByMemberId(ctx, userId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	for _, v := range memberWallet {
		coinInfo, err := d.marketRpc.FindCoinInfo(ctx, &mclient.MarketReq{
			Unit: v.CoinName,
		})
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		if coinInfo.Unit == "USDT" {
			coinInfo.CnyRate = 7
			coinInfo.UsdRate = 1
		} else {
			coinInfo.UsdRate = 20000
			coinInfo.CnyRate = op.MulFloor(7, coinInfo.UsdRate, 10)
		}
		list = append(list, v.Copy(coinInfo))

	}
	return list, nil
	//查询 币种详情

}

func NewMemberWalletDomain(db *msdb.MsDB, marketRpc mclient.Market) *MemberWalletDomain {
	return &MemberWalletDomain{
		memberWalletRepo: dao.NewMemberWalletDao(db),
		transaction:      tran.NewTransaction(db.Conn),
		marketRpc:        marketRpc,
	}
}
