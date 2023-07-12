package repo

import (
	"context"
	"mscoin-common/msdb"
	"ucenter/internal/model"
)

type MemberWalletRepo interface {
	Save(ctx context.Context, mw *model.MemberWallet) error
	FindByIdAndCoinName(ctx context.Context, memId int64, coinName string) (mw *model.MemberWallet, err error)
	UpdateFreeze(ctx context.Context, conn msdb.DbConn, memberId int64, symbol string, money float64) error
}
