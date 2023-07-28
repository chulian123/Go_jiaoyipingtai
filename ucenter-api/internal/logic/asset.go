package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/ucenter/types/asset"
	"mscoin-common/pages"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type Asset struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Asset {
	return &Asset{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Asset) FindWalletBySymbol(req *types.AssetReq) (*types.MemberWallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	memberWallet, err := l.svcCtx.UCAssetRpc.FindWalletBySymbol(ctx, &asset.AssetReq{
		CoinName: req.CoinName,
		UserId:   value,
	})
	if err != nil {
		return nil, err
	}
	resp := &types.MemberWallet{}
	if err := copier.Copy(resp, memberWallet); err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *Asset) FindWallet(req *types.AssetReq) ([]*types.MemberWallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	memberWalletResp, err := l.svcCtx.UCAssetRpc.FindWallet(ctx, &asset.AssetReq{
		UserId: value,
	})
	if err != nil {
		return nil, err
	}
	var resp []*types.MemberWallet
	if err := copier.Copy(&resp, memberWalletResp.List); err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *Asset) ResetAddress(req *types.AssetReq) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	_, err := l.svcCtx.UCAssetRpc.ResetAddress(ctx, &asset.AssetReq{
		UserId:   value,
		CoinName: req.Unit,
	})
	if err != nil {
		return "", err
	}
	return "", nil
}

func (l *Asset) FindTransaction(req *types.AssetReq) (*pages.PageResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	value := l.ctx.Value("userId").(int64)
	resp, err := l.svcCtx.UCAssetRpc.FindTransaction(ctx, &asset.AssetReq{
		UserId:    value,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		PageNo:    int64(req.PageNo),
		PageSize:  int64(req.PageSize),
		Symbol:    req.Symbol,
		Type:      req.Type,
	})
	if err != nil {
		return nil, err
	}
	total := resp.Total
	b := make([]any, len(resp.List))
	for i, v := range resp.List {
		b[i] = v
	}
	return pages.New(b, int64(req.PageNo), int64(req.PageSize), total), nil
}
