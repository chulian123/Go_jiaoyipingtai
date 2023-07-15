package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	memberDomain       *domain.MemberDomain
	memberWalletDomain *domain.MemberWalletDomain
}

func (l *AssetLogic) FindWalletBySymbol(req *asset.AssetReq) (*asset.MemberWallet, error) {
	//通过market rpc 进行coin表的查询 coin信息
	//通过钱包 查询对应币的钱包信息  coin_id  user_id 查询用户的钱包信息 组装信息
	coinInfo, err := l.svcCtx.MarketRpc.FindCoinInfo(l.ctx, &market.MarketReq{
		Unit: req.CoinName,
	})
	if err != nil {
		return nil, err
	}
	memberWalletCoin, err := l.memberWalletDomain.FindWalletBySymbol(l.ctx, req.UserId, req.CoinName, coinInfo)
	if err != nil {
		return nil, err
	}
	resp := &asset.MemberWallet{}
	copier.Copy(resp, memberWalletCoin)
	return resp, nil
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		memberDomain:       domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain: domain.NewMemberWalletDomain(svcCtx.Db),
	}
}
