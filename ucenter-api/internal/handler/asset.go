package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	common "mscoin-common"
	"mscoin-common/tools"
	"net/http"
	"ucenter-api/internal/logic"
	"ucenter-api/internal/svc"
	"ucenter-api/internal/types"
)

type AssetHandler struct {
	svcCtx *svc.ServiceContext
}

func NewAssetHandler(svcCtx *svc.ServiceContext) *AssetHandler {
	return &AssetHandler{
		svcCtx: svcCtx,
	}
}

func (h *AssetHandler) FindWalletBySymbol(w http.ResponseWriter, r *http.Request) {
	var req types.AssetReq
	if err := httpx.ParsePath(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	resp, err := l.FindWalletBySymbol(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *AssetHandler) FindWallet(w http.ResponseWriter, r *http.Request) {
	var req = types.AssetReq{}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewAssetLogic(r.Context(), h.svcCtx)
	resp, err := l.FindWallet(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
