package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/withdraw"
	"ucenter/internal/domain"
	"ucenter/internal/svc"
)

type WithdrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberAddressDomain *domain.MemberAddressDomain
}

func (l *WithdrawLogic) FindAddressByCoinId(req *withdraw.WithdrawReq) (*withdraw.AddressSimpleList, error) {
	ctx := context.Background()
	userId := req.UserId
	coinId := req.CoinId
	memberAddresses, err := l.memberAddressDomain.FindAddressByCoinId(ctx, userId, coinId)
	if err != nil {
		return nil, err
	}
	var list []*withdraw.AddressSimple
	copier.Copy(&list, memberAddresses)
	return &withdraw.AddressSimpleList{
		List: list,
	}, nil
}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		ctx:                 ctx,
		svcCtx:              svcCtx,
		Logger:              logx.WithContext(ctx),
		memberAddressDomain: domain.NewMemberAddressDomain(svcCtx.Db),
	}
}
