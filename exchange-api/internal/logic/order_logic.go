package logic

import (
	"context"
	"errors"
	"exchange-api/internal/svc"
	"exchange-api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/exchange/types/order"
	"mscoin-common/pages"
	"time"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (l *OrderLogic) History(req *types.ExchangeReq) (*pages.PageResult, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	userId := l.ctx.Value("userId").(int64)
	symbol := req.Symbol
	orderRes, err := l.svcCtx.OrderRpc.FindOrderHistory(ctx, &order.OrderReq{
		Symbol:   symbol,
		Page:     req.PageNo,
		PageSize: req.PageSize,
		UserId:   userId,
	})
	if err != nil {
		return nil, err
	}
	list := orderRes.List
	b := make([]any, len(list))
	for i := range list {
		b[i] = list[i]
	}
	return pages.New(b, req.PageNo, req.PageSize, orderRes.Total), nil
}

func (l *OrderLogic) Current(req *types.ExchangeReq) (*pages.PageResult, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 10*time.Second)
	defer cancel()
	userId := l.ctx.Value("userId").(int64)
	symbol := req.Symbol
	orderRes, err := l.svcCtx.OrderRpc.FindOrderCurrent(ctx, &order.OrderReq{
		Symbol:   symbol,
		Page:     req.PageNo,
		PageSize: req.PageSize,
		UserId:   userId,
	})
	if err != nil {
		return nil, err
	}
	list := orderRes.List
	b := make([]any, len(list))
	for i := range list {
		b[i] = list[i]
	}
	return pages.New(b, req.PageNo, req.PageSize, orderRes.Total), nil
}

func (l *OrderLogic) AddOrder(req *types.ExchangeReq) (string, error) {
	//调用exchange rpc 完成addOrder 功能
	value := l.ctx.Value("userId").(int64)
	if !req.OrderValid() {
		return "", errors.New("参数传递错误")
	}
	orderRes, err := l.svcCtx.OrderRpc.Add(l.ctx, &order.OrderReq{
		Symbol:    req.Symbol,
		UserId:    value,
		Direction: req.Direction,
		Type:      req.Type,
		Price:     req.Price,
		Amount:    req.Amount,
	})
	if err != nil {
		return "", err
	}
	return orderRes.OrderId, nil
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
