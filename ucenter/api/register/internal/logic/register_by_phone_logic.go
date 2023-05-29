package logic

import (
	"context"

	"ucenter/api/register/internal/svc"
	"ucenter/api/types/register"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterByPhoneLogic {
	return &RegisterByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterByPhoneLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line

	return &register.RegRes{}, nil
}
