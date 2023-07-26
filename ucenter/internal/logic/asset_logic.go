package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"mscoin-common/base"
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

func (l *AssetLogic) FindWallet(req *asset.AssetReq) (*asset.MemberWalletList, error) {
	//根据用户id来查询用户钱包 循坏钱包信息 根据币种 查询币种详情
	wallet, err := l.memberWalletDomain.FindWallet(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*asset.MemberWallet
	copier.Copy(&list, wallet)
	return &asset.MemberWalletList{
		List: list,
	}, nil
}

func (l *AssetLogic) ResetAddress(req *asset.AssetReq) (*asset.AssetResp, error) {
	//查询用户钱包 查询用户钱包地址 地址为空就要生成新的
	memberWallet, err := l.memberWalletDomain.FindWalletByMemIdAndCoin(l.ctx, req.UserId, req.CoinName)
	if err != nil {
		return nil, err
	}
	//判断比特币的
	if req.CoinName == "BTC" {
		if memberWallet.Address == "" {
			wallet, err := base.NewWallet()
			if err != nil {
				logx.Info("生成NewWallet失败!")
				return nil, err
			}
			address := wallet.GetTestAddress()
			privatekey := wallet.GetPriKey()
			memberWallet.AddressPrivateKey = privatekey
			memberWallet.Address = string(address)
			//更新钱包信息
			err = l.memberWalletDomain.UpdateAddress(l.ctx, memberWallet)
			if err != nil {
				logx.Info("更新钱包信息失败!")
				return nil, err
			}
		}
	}
	return &asset.AssetResp{}, nil
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		memberDomain:       domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain: domain.NewMemberWalletDomain(svcCtx.Db, svcCtx.MarketRpc, svcCtx.Cache),
	}
}
