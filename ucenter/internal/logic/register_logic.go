package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"mscoin-common/tools"
	"time"
	"ucenter/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "REGISTER::"

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	logx.Info("ucenter Register By Phone call服务调用成功 ...")
	//1.校验前端的人际验证是否通过
	//请求参数为JSON对象
	//https://www.vaptcha.com/document/install.html#%E6%9C%8D%E5%8A%A1%E7%AB%AF%E4%BA%8C%E6%AC%A1%E9%AA%8C%E8%AF%81
	isVerify := l.CaptchaDomain.Verify(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip)
	if !isVerify {
		return nil, errors.New("人机校验不通过(来自 Ucenter RegisterByPhone 报错)")
	}
	logx.Info("人机校验通过 ...")

	//2.校验验证码
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	redisValue := "" //设置验证码为空数值，
	err := l.svcCtx.Cache.GetCtx(ctx, RegisterCacheKey+in.Phone, redisValue)
	if err != nil {
		return nil, errors.New("redis验证码获取错误")
	}
	//如果上面没有出错，就把存在redis里面的验证码获取到，然后和前端的传入的验证码校验
	if in.Code != redisValue {
		return nil, errors.New("验证码输入不正确")
	}

	//3.验证码通过后 进行注册步骤

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
