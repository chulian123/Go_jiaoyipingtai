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
	memberDomain            *domain.MemberDomain
	memberWalletDomain      *domain.MemberWalletDomain
	memberTransactionDomain *domain.MemberTransactionDomain
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
	//根据用户id查询用户的钱包 循环钱包信息 根据币种 查询币种详情
	memberWalletCoins, err := l.memberWalletDomain.FindWallet(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var list []*asset.MemberWallet
	copier.Copy(&list, memberWalletCoins)
	return &asset.MemberWalletList{
		List: list,
	}, nil
}

func (l *AssetLogic) ResetAddress(req *asset.AssetReq) (*asset.AssetResp, error) {
	//查询用户的钱包 检查address是否为空 如果未空 生成地址 进行更新
	memberWallet, err := l.memberWalletDomain.FindWalletByMemIdAndCoin(l.ctx, req.UserId, req.CoinName)
	if err != nil {
		return nil, err
	}
	if req.CoinName == "BTC" {
		if memberWallet.Address == "" {
			wallet, err := base.NewWallet()
			if err != nil {
				return nil, err
			}
			address := wallet.GetTestAddress()
			priKey := wallet.GetPriKey()
			memberWallet.AddressPrivateKey = priKey
			memberWallet.Address = string(address)
			err = l.memberWalletDomain.UpdateAddress(l.ctx, memberWallet)
			if err != nil {
				return nil, err
			}
		}
	}
	return &asset.AssetResp{}, nil
}

func (l *AssetLogic) FindTransaction(req *asset.AssetReq) (*asset.MemberTransactionList, error) {
	//查询所有的充值记录 分页查询
	memberTransactionVos, total, err := l.memberTransactionDomain.FindTransaction(
		l.ctx,
		req.UserId,
		req.PageNo,
		req.PageSize,
		req.Symbol,
		req.StartTime,
		req.EndTime,
		req.Type,
	)
	if err != nil {
		return nil, err
	}
	var list []*asset.MemberTransaction
	copier.Copy(&list, memberTransactionVos)
	return &asset.MemberTransactionList{
		List:  list,
		Total: total,
	}, nil
}

func (l *AssetLogic) GetAddress(req *asset.AssetReq) (*asset.AddressList, error) {
	addressList, err := l.memberWalletDomain.GetAllAddress(l.ctx, req.CoinName)
	if err != nil {
		return nil, err
	}
	return &asset.AddressList{
		List: addressList,
	}, nil
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssetLogic {
	return &AssetLogic{
		ctx:                     ctx,
		svcCtx:                  svcCtx,
		Logger:                  logx.WithContext(ctx),
		memberDomain:            domain.NewMemberDomain(svcCtx.Db),
		memberWalletDomain:      domain.NewMemberWalletDomain(svcCtx.Db, svcCtx.MarketRpc, svcCtx.Cache),
		memberTransactionDomain: domain.NewMemberTransactionDomain(svcCtx.Db),
	}
}
