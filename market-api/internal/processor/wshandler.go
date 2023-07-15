package processor

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/model"
	"market-api/internal/ws"
)

type WebsocketHandler struct {
	wsServer *ws.WebsocketServer
}

func (w *WebsocketHandler) HandleTrade(symbol string, data []byte) {
	//TODO implement me
	panic("implement me")
}

func (w *WebsocketHandler) HandleKLine(symbol string, kline *model.Kline, thumbMap map[string]*market.CoinThumb) {
	logx.Info("================WebsocketHandler Start=======================")
	logx.Info("symbol:", symbol)
	thumb := thumbMap[symbol]
	if thumb == nil {
		thumb = kline.InitCoinThumb(symbol)
	}
	coinThumb := kline.ToCoinThumb(symbol, thumb)
	result := &model.CoinThumb{}
	copier.Copy(result, coinThumb)
	marshal, _ := json.Marshal(result)
	w.wsServer.BroadcastToNamespace("/", "/topic/market/thumb", string(marshal))

	bytes, _ := json.Marshal(kline)
	w.wsServer.BroadcastToNamespace("/", "/topic/market/kline/"+symbol, string(bytes))

	logx.Info("================WebsocketHandler End=======================")
}

func NewWebsocketHandler(wsServer *ws.WebsocketServer) *WebsocketHandler {
	return &WebsocketHandler{
		wsServer: wsServer,
	}
}
