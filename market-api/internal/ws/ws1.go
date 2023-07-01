package ws

//集成WebsocketIO
//
//const ROOM = "market"
//
//// path表示WebSocket的路径，server表示WebSocket服务器实例。
//type WebSocketServer struct {
//	path   string
//	server *socketio.Server
//}
//
//func (ws *WebSocketServer) Start() {
//	ws.server.Serve()
//}
//
//func (ws *WebSocketServer) Stop() {
//	ws.server.Close()
//}
//
//var allowOriginFunc = func(r *http.Request) bool {
//	return true
//}
//
//func NewWebSocketServer(path string) *WebSocketServer {
//	//解决跨域问题
//	server := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{
//			&polling.Transport{
//				CheckOrigin: allowOriginFunc,
//			},
//			&websocket.Transport{
//				CheckOrigin: allowOriginFunc,
//			},
//		},
//	})
//	//WebSocket服务器的连接事件进行处理
//	server.OnConnect("/", func(s socketio.Conn) error {
//		s.SetContext("")
//		s.Join(ROOM)
//		logx.Info("connected:", s.ID())
//		return nil
//	})
//	return &WebSocketServer{
//		path:   path,
//		server: server,
//	}
//}
//
//// BroadcastToNamespace  "/" "/topic/market/thumb"
//// WebSocket服务器实例中的方法，用于向指定房间（ROOM）广播事件
//func (w *WebSocketServer) BroadcastToNamespace(path string, event string, data any) {
//	//后台异步执行广播操作
//	go func() {
//		w.server.BroadcastToRoom(path, ROOM, event, data) //将事件和数据广播到指定的命名空间（path）中的指定房间（ROOM），使得连接到该房间的客户端都能收到该事件和数据
//	}()
//	//以实现非阻塞的异步广播，避免阻塞当前线程的执行
//}
//
//func (ws *WebSocketServer) ServerHandler(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		//所有的请求都会经过这里
//		path := r.URL.Path
//		logx.Info("=======ServerHandler", path)
//		if strings.HasPrefix(path, ws.path) { //strings.HasPrefix用来检测字符串是否以指定的前缀开头。
//			//如果为真，进行我们的处理
//			logx.Info("执行为真")
//			ws.server.ServeHTTP(w, r)
//		} else {
//			logx.Info("执行为假")
//			next.ServeHTTP(w, r)
//		}
//	})
//}
