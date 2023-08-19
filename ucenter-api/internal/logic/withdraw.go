package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"grpc-common/ucenter/types/asset"
	"grpc-common/ucenter/types/member"
	"grpc-common/ucenter/types/withdraw"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type Withdraw struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (w *Withdraw) QueryWithdrawCoin(req *types.WithdrawReq) ([]*types.WithdrawWalletInfo, error) {
	//1.查询所有的coin信息
	coinList, err := w.svcCtx.MarketRpc.FindAllCoin(w.ctx, &market.MarketReq{})
	if err != nil {
		return nil, err
	}
	coinMap := make(map[string]*market.Coin)
	for _, v := range coinList.List {
		coinMap[v.Unit] = v //传入对应单位
	}
	//2.根据用户id 查询用户信息
	userId := w.ctx.Value("userId").(int64)
	walletList, err := w.svcCtx.UCAssetRpc.FindWallet(w.ctx, &asset.AssetReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	//3.组装数据[]*types.WithdrawWalletInfo
	wwlist := make([]*types.WithdrawWalletInfo, len(walletList.List))
	for i, wallet := range walletList.List {
		var ww types.WithdrawWalletInfo
		coin := coinMap[wallet.Coin.Unit]
		ww.Balance = wallet.Balance
		ww.WithdrawScale = int(coin.WithdrawScale)
		ww.MaxTxFee = coin.MaxTxFee
		ww.MinTxFee = coin.MinTxFee
		ww.MaxAmount = coin.MaxWithdrawAmount
		ww.MinAmount = coin.MinWithdrawAmount
		ww.Name = coin.GetName()
		ww.NameCn = coin.NameCn
		ww.Threshold = coin.WithdrawThreshold
		ww.Unit = coin.Unit
		ww.AccountType = int(coin.AccountType)
		if coin.CanAutoWithdraw == 0 {
			ww.CanAutoWithdraw = "true"
		} else {
			ww.CanAutoWithdraw = "false"
		}
		//提币地址的赋值
		addressSimpleList, err := w.svcCtx.UCWithdrawRpc.FindAddressByCoinId(w.ctx, &withdraw.WithdrawReq{
			UserId: userId,
			CoinId: int64(coin.Id),
		})
		if err != nil {
			return nil, err
		}
		var addressList []types.AddressSimple
		copier.Copy(&addressList, addressSimpleList.List)
		ww.Addresses = addressList
		wwlist[i] = &ww
	}
	return wwlist, nil

}

func (w *Withdraw) SengCode(t *types.WithdrawReq) (string, error) {
	//1.根据当前登录用户id 查询用户消息 获取到手机号
	userId := w.ctx.Value("userId").(int64)
	memberInfo, err := w.svcCtx.UCMemberRpc.FindMemberById(w.ctx, &member.MemberReq{MemberId: userId})
	if err != nil {
		return "", err
	}
	phone := memberInfo.MobilePhone
	//2.根据注册手机号发送验证码
	_, err = w.svcCtx.UCWithdrawRpc.SendCode(w.ctx, &withdraw.WithdrawReq{
		Phone: phone,
	})
	if err != nil {
		return "", err
	}
	return "success", nil

}

func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Withdraw {
	return &Withdraw{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
