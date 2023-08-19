package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type WithdrawHandler struct {
	svcCtx *svc.ServiceContext
}

func (h *WithdrawHandler) QueryWithdrawCoin(w http.ResponseWriter, r *http.Request) {
	var req types.WithdrawReq
	l := logic.NewWithdrawLogic(r.Context(), h.svcCtx)
	resp, err := l.QueryWithdrawCoin(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *WithdrawHandler) SengCode(w http.ResponseWriter, r *http.Request) {
	var req types.WithdrawReq
	l := logic.NewWithdrawLogic(r.Context(), h.svcCtx)
	resp, err := l.SengCode(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func NewWithdrawHandler(svcCtx *svc.ServiceContext) *WithdrawHandler {
	return &WithdrawHandler{svcCtx}
}
