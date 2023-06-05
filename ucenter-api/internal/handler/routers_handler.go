package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	//封装一下的url
	//如果需要加中间件 怎么办
	//注册url
	register := NewRegisterHandler(serverCtx)
	registerGroup := r.Group()
	registerGroup.Post("/uc/register/phone", register.Register)
	registerGroup.Post("/uc/mobile/code", register.SendCode)

	//登录路由
	login := NewLoginHandler(serverCtx)
	loginGroup := r.Group()
	loginGroup.Post("/uc/login", login.Login)
}
