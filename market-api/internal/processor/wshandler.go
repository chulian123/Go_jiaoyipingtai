package processor

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/model"
	"market-api/internal/ws"
)

type WebSocketHandler struct {
	wsServer *ws.WebsocketServer
}

func (w *WebSocketHandler) HandleTrade(symbol string, data []byte) {
	//TODO implement me
	panic("implement me")
}

func (w *WebSocketHandler) HandleKLine(symbol string, kline *model.Kline, thumpMap map[string]*market.CoinThumb) {
	logx.Info("================WebsocketHandler Start=======================")
	logx.Info("symbol:", symbol)
	thumb := thumpMap[symbol]
	if thumb == nil {
		kline.InitCoinThumb(symbol)
	}
	coinThumb := kline.ToCoinThumb(symbol, thumb)
	marshal, _ := json.Marshal(coinThumb)
	w.wsServer.BroadcastToNamespace("/", "/topic/market/thumb", string(marshal))

	logx.Info("================WebsocketHandler end=======================")
}

func NewWebSocketHandler(wsServer *ws.WebsocketServer) *WebSocketHandler {
	return &WebSocketHandler{
		wsServer: wsServer,
	}
}
