package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"market-api/internal/logic"
	"market-api/internal/svc"
	"market-api/internal/types"
	common "mscoin-common"
	"mscoin-common/tools"
	"net/http"
)

//注册的handlər

type MarketHandler struct {
	svcCtx *svc.ServiceContext
}

func NewMarketHandler(svcCtx *svc.ServiceContext) *MarketHandler {
	return &MarketHandler{
		svcCtx: svcCtx,
	}
}

// SymbolThumbTrend 获取币种行情
func (h *MarketHandler) SymbolThumbTrend(w http.ResponseWriter, r *http.Request) {
	var req = &types.MarketReq{}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolThumbTrend(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// SymbolThumb 获取币种行情
func (h *MarketHandler) SymbolThumb(w http.ResponseWriter, r *http.Request) {
	var req = &types.MarketReq{}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolThumb(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

// SymbolInfo 交易币详情
func (h *MarketHandler) SymbolInfo(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil { //ParseForm parses the form request.
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := common.NewResult()
	//获取一下ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.SymbolInfo(req)
	result := newResult.Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)

}

func (h *MarketHandler) CoinInfo(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.CoinInfo(&req)
	result := common.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}

func (h *MarketHandler) History(w http.ResponseWriter, r *http.Request) {
	var req types.MarketReq
	if err := httpx.ParseForm(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	ip := tools.GetRemoteClientIp(r)
	req.Ip = ip
	l := logic.NewMarketLogic(r.Context(), h.svcCtx)
	resp, err := l.History(&req)
	result := common.NewResult().Deal(resp.List, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
