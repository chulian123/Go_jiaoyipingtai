package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"grpc-common/ucenter/types/login"
	"mscoin-common/tools"
	"time"
	"ucenter/internal/domain"
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
		MemberDomain:  domain.NewMemberDomain(svcCtx.Db),
	}
}

// Login {code:111,msg:xxx}
func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginRes, error) {
	//1. 先校验人机是否通过
	isVerify := l.CaptchaDomain.Verify(
		in.Captcha.Server,
		l.svcCtx.Config.Captcha.Vid,
		l.svcCtx.Config.Captcha.Key,
		in.Captcha.Token,
		2,
		in.Ip)
	if !isVerify {
		return nil, errors.New("人机校验不通过")
	}
	//2. 校验密码
	member, err := l.MemberDomain.FindByPhone(context.Background(), in.GetUsername())
	if err != nil {
		logx.Error(err)
		return nil, errors.New("登录失败")
	}
	if member == nil {
		return nil, errors.New("此用户未注册")
	}
	password := member.Password
	salt := member.Salt
	verify := tools.Verify(in.Password, salt, password, nil)
	if !verify {
		return nil, errors.New("密码不正确")
	}
	//3. 登录成功，生成token，提供给前端，前端调用传递token，我们进行token认证即可
	//jwt的技术 A.B.C
	key := l.svcCtx.Config.JWT.AccessSecret
	expire := l.svcCtx.Config.JWT.AccessExpire
	token, err := l.getJwtToken(key, time.Now().Unix(), expire, member.Id)
	if err != nil {
		return nil, errors.New("token生成错误")
	}
	//返回登录所需信息
	loginCount := member.LoginCount + 1
	go func() {
		l.MemberDomain.UpdateLoginCount(context.Background(), member.Id, 1)
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
		LoginCount:    int32(loginCount),
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
