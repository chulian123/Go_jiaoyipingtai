package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"mscoin-common/tools"
	"time"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER::"

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	// todo: add your logic here and delete this line
	logx.Info("ucenter Register By Phone call ...")
	return &register.RegRes{}, nil
}

func (l *RegisterLogic) SendCode(req *register.CodeReq) (*register.NoRes, error) {
	//验证码逻辑：
	//* 收到手机号和国家标识
	//* 生成验证码
	//* 根据对应的国家和手机号调用对应的短信平台发送验证码
	//* 将验证码存入redis，过期时间5分钟
	//* 返回成功
	code := tools.Rand4Code() //生成4位验证码
	//假设调用短信平台服务成功
	go func() {
		logx.Info("调用短信平台成功")
	}()
	logx.Infof("验证码为: %s \n", code)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 15*time.Minute) //存入redis，15分钟后过期
	if err != nil {
		return nil, errors.New("验证码发送失败,验证码存入cashed失败")

	}
	return &register.NoRes{}, nil

}
