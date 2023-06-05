package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"grpc-common/ucenter/types/login"
	"mscoin-common/tools"
	"time"

	"ucenter/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db), //传入数据库链接
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {
	logx.Info("ucenter Login By Phone call服务调用成功 ...")
	//1.校验前端的人际验证是否通过
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
	//2,密码校验
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	member, err := l.MemberDomain.FindByPhone(ctx, in.GetUsername())
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败")
	}
	if member == nil { //如果查询回来的登录用户为空
		logx.Error(err)
		return nil, errors.New("该用户未注册")
	}
	password := member.Password
	salt := member.Salt
	verify := tools.Verify(in.Password, salt, password, nil) //解密密码
	if !verify {
		return nil, errors.New("密码不正确")
	}
	//3，登录成功后 生成token提供给前端，前端调用token来传递， 我们进行token认证
	key := l.svcCtx.Config.JWT.AccessSecret        //获取加密的钥匙窜
	expireTime := l.svcCtx.Config.JWT.AccessExpire //token的过期时间
	token, err := l.getJwtToken(key, time.Now().Unix(), expireTime, member.Id)
	if err != nil {
		logx.Error(err)
		return nil, errors.New("Token生成错误")
	}
	//返回登录信息
	loginCount := member.LoginCount + 1
	go func() {
		l.MemberDomain.UpDateLoginCount(context.Background(), member.Id, 1)
	}()
	return &login.LoginRes{
		Token:         token,
		Id:            member.Id,
		Username:      member.Username,
		MemberLevel:   member.MemberLevelStr(),
		MemberRate:    member.MemberRate(),
		RealName:      member.RealName,
		Country:       member.Country,
		Avatar:        member.Avatar,
		PromotionCode: member.PromotionCode,
		SuperPartner:  member.SuperPartner,
		LoginCount:    int32(loginCount), //登录次数
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
