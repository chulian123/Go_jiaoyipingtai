package logic

//注册业务的逻辑代码

import (
	"context"
	"grpc-common/ucenter/types/register"
	"time"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	logx.Info("api Register")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second) //设置5秒超时时间 防止链接长时间占用资源
	defer cancelFunc()
	_, err = l.svcCtx.UCRegisterRpc.RegisterByPhone(ctx, &register.RegReq{})
	if err != nil {
		return nil, err //如果注册失败就返回
	}
	return
}

func (l *RegisterLogic) SendCode(req *types.CodeRequest) (resp *types.CodeResponse, err error) {
	logx.Info("api SendCode")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second) //设置5秒超时时间 防止链接长时间占用资源
	defer cancelFunc()
	_, err = l.svcCtx.UCRegisterRpc.SendCode(ctx, &register.CodeReq{
		Phone:   req.Phone,
		Country: req.Country,
	})
	if err != nil {
		return nil, err //如果注册失败就返回
	}
	return
}
