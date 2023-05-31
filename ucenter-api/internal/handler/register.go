package handler

//注册的handlər

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"

	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}

func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request

	if err := httpx.Parse(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	//对前端的人机识别参数进行校验
	newResult := common.NewResult()
	if req.Captcha == nil {
		httpx.OkJsonCtx(r.Context(), w, newResult.Deal(nil, errors.New("人机校验不通过！(来自 Ucenter-api RegisterByPhone 报错)")))
		return
	}
	//获取ip
	req.Ip = tools.GetRemoteClientIp(r)

	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.Register(&req)

	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)

}

func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.SendCode(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)

}
