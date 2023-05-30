package handler

import (
	"ucenter-api/internal/svc"
)

func RegisterHandlers(r *Routers, serverCtx *svc.ServiceContext) {
	//封装一下的url
	//如果需要加中间件 怎么办
	register := NewRegisterHandler(serverCtx)
	registerRouters := r.Group()
	registerRouters.Get("/uc/register/phone", register.Register)

}
